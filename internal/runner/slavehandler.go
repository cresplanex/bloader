package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"
	"google.golang.org/protobuf/proto"

	"github.com/cresplanex/bloader/internal/logger"
)

// SlaveRequestHandler is a struct that holds the response handler information.
type SlaveRequestHandler struct {
	// resChan is a channel.
	resChan <-chan *pb.ReceiveChanelConnectResponse
	// cli is a client.
	cli pb.BloaderSlaveServiceClient
	// chunkSize is an integer.
	chunkSize int
	// receiveTermChan is a channel for receiving term.
	receiveTermChan <-chan ReceiveTermType
	// dataBufferMap is a map.
	dataBufferMap map[string]*bytes.Buffer
}

// DefaultChunkSize is an integer.
const DefaultChunkSize = 1024

// NewSlaveRequestHandler creates a new ResponseHandler.
func NewSlaveRequestHandler(
	resChan <-chan *pb.ReceiveChanelConnectResponse,
	cli pb.BloaderSlaveServiceClient,
	termChan <-chan ReceiveTermType,
) *SlaveRequestHandler {
	return &SlaveRequestHandler{
		resChan:         resChan,
		cli:             cli,
		chunkSize:       DefaultChunkSize,
		receiveTermChan: termChan,
		dataBufferMap:   make(map[string]*bytes.Buffer),
	}
}

// HandleResponse handles the response.
func (rh *SlaveRequestHandler) HandleResponse(
	ctx context.Context,
	log logger.Logger,
	tmplFactor TmplFactor,
	authFactor AuthenticatorFactor,
	targetFactor TargetFactor,
	store Store,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case termType := <-rh.receiveTermChan:
			switch termType {
			case ReceiveTermTypeReceiveTermTypeEOF:
				return nil
			case ReceiveTermTypeReceiveTermTypeResponseReceiveError:
				return fmt.Errorf("response receive error")
			case ReceiveTermTypeReceiveTermTypeStreamContextDone:
				return nil
			case ReceiveTermTypeReceiveTermTypeDisconnected:
				return nil
			default:
				return fmt.Errorf("unknown term type: %v", termType)
			}
		case res := <-rh.resChan:
			log.Debug(ctx, "Received response: %v",
				logger.Value("response", res))
			if res == nil {
				return fmt.Errorf("response is nil")
			}
			switch res.RequestType {
			case pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_LOADER:
				loaderResourceReq := res.GetLoaderResourceRequest()
				stream, err := rh.cli.SendLoader(ctx)
				if err != nil {
					return fmt.Errorf("failed to send loader: %w", err)
				}
				tmplStr, err := tmplFactor.TmplFactorize(ctx, loaderResourceReq.LoaderId)
				if err != nil {
					return fmt.Errorf("failed to factorize template: %w", err)
				}
				buffer := []byte(tmplStr)
				for i := 0; i < len(buffer); i += rh.chunkSize {
					end := i + rh.chunkSize
					if end > len(buffer) {
						end = len(buffer)
					}
					if err := stream.Send(&pb.SendLoaderRequest{
						RequestId:   res.RequestId,
						LoaderId:    loaderResourceReq.LoaderId,
						Content:     buffer[i:end],
						IsLastChunk: end == len(buffer),
					}); err != nil {
						return fmt.Errorf("failed to send loader request: %w", err)
					}
				}
				_, err = stream.CloseAndRecv()
				if err != nil {
					return fmt.Errorf("failed to receive loader response: %w", err)
				}
				log.Info(ctx, "Sent loader: %v",
					logger.Value("loader_id", loaderResourceReq.LoaderId))
			case pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_AUTH:
				authResourceReq := res.GetAuthResourceRequest()
				auth, err := authFactor.Factorize(ctx, authResourceReq.AuthId, authResourceReq.IsDefault)
				if err != nil {
					return fmt.Errorf("failed to factorize auth: %w", err)
				}
				_, err = rh.cli.SendAuth(ctx, &pb.SendAuthRequest{
					RequestId: res.RequestId,
					AuthId:    authResourceReq.AuthId,
					Auth:      auth.GetAuthValue(),
					IsDefault: authFactor.IsDefault(authResourceReq.AuthId),
				})
				if err != nil {
					return fmt.Errorf("failed to send auth: %w", err)
				}
				log.Info(ctx, "Sent auth: %v",
					logger.Value("auth_id", authResourceReq.AuthId))
			case pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_STORE:
				storeResourceReq := res.GetStoreResourceRequest()
				buffer, ok := rh.dataBufferMap[res.RequestId]
				if !ok {
					buffer = &bytes.Buffer{}
					rh.dataBufferMap[res.RequestId] = buffer
				}
				if _, err := buffer.Write(storeResourceReq.Data); err != nil {
					return fmt.Errorf("failed to write import store data: %w", err)
				}
				if storeResourceReq.IsLastChunk {
					byteData := buffer.Bytes()
					delete(rh.dataBufferMap, res.RequestId)
					var dataList pb.StoreImportRequestList
					if err := proto.Unmarshal(byteData, &dataList); err != nil {
						return fmt.Errorf("failed to unmarshal import store data list: %w", err)
					}

					storeData := make([]ValidStoreImportData, len(dataList.Data))
					for i, data := range dataList.Data {
						storeData[i] = ValidStoreImportData{
							BucketID: data.BucketId,
							StoreKey: data.StoreKey,
						}
						if data.Encryption != nil {
							storeData[i].Encrypt = ValidCredentialEncryptConfig{
								Enabled:   data.Encryption.Enabled,
								EncryptID: data.Encryption.EncryptId,
							}
						}
					}
					strData := make([]*pb.StoreExportData, 0, len(dataList.Data))
					if err := store.Import(
						ctx,
						storeData,
						func(_ context.Context, data ValidStoreImportData, val any, valBytes []byte) error {
							var err error
							if valBytes == nil {
								valBytes, err = json.Marshal(val)
								if err != nil {
									return fmt.Errorf("failed to marshal store data: %w", err)
								}
							}
							strData = append(strData, &pb.StoreExportData{
								BucketId: data.BucketID,
								StoreKey: data.StoreKey,
								Data:     valBytes,
							})
							return nil
						},
					); err != nil {
						return fmt.Errorf("failed to import store data: %w", err)
					}

					stream, err := rh.cli.SendStoreData(ctx)
					if err != nil {
						return fmt.Errorf("failed to send store data: %w", err)
					}
					strDataList := &pb.StoreExportDataList{
						Data: strData,
					}
					b, err := proto.Marshal(strDataList)
					if err != nil {
						return fmt.Errorf("failed to marshal store data list: %w", err)
					}
					for i := 0; i < len(b); i += rh.chunkSize {
						end := i + rh.chunkSize
						if end > len(b) {
							end = len(b)
						}
						if err := stream.Send(&pb.SendStoreDataRequest{
							RequestId:   res.RequestId,
							Data:        b[i:end],
							IsLastChunk: end == len(b),
						}); err != nil {
							return fmt.Errorf("failed to send store data request: %w", err)
						}
					}
					_, err = stream.CloseAndRecv()
					if err != nil {
						return fmt.Errorf("failed to receive store data response: %w", err)
					}
					log.Info(ctx, "Sent store: %v",
						logger.Value("store_data", strData))
				}
			case pb.RequestType_REQUEST_TYPE_STORE:
				storeReq := res.GetStore()
				buffer, ok := rh.dataBufferMap[res.RequestId]
				if !ok {
					buffer = &bytes.Buffer{}
					rh.dataBufferMap[res.RequestId] = buffer
				}
				if _, err := buffer.Write(storeReq.Data); err != nil {
					return fmt.Errorf("failed to write store data: %w", err)
				}
				if storeReq.IsLastChunk {
					byteData := buffer.Bytes()
					delete(rh.dataBufferMap, res.RequestId)
					var dataList pb.StoreDataList
					if err := proto.Unmarshal(byteData, &dataList); err != nil {
						return fmt.Errorf("failed to unmarshal store data list: %w", err)
					}

					storeData := make([]ValidStoreValueData, len(dataList.Data))
					for i, data := range dataList.Data {
						var val any
						if err := json.Unmarshal(data.Data, &val); err != nil {
							return fmt.Errorf("failed to unmarshal store data: %w", err)
						}
						storeData[i] = ValidStoreValueData{
							BucketID: data.BucketId,
							Key:      data.StoreKey,
							Value:    val,
							Encrypt: ValidCredentialEncryptConfig{
								Enabled:   data.Encryption.Enabled,
								EncryptID: data.Encryption.EncryptId,
							},
						}
					}

					if err := store.Store(ctx, storeData, nil); err != nil {
						return fmt.Errorf("failed to store data: %w", err)
					}
					if _, err := rh.cli.SendStoreOk(ctx, &pb.SendStoreOkRequest{
						RequestId: res.RequestId,
					}); err != nil {
						return fmt.Errorf("failed to send store ok: %w", err)
					}
					log.Info(ctx, "Stored data: %v",
						logger.Value("store_data", storeData))
				}
			case pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_TARGET:
				targetResourceReq := res.GetTargetResourceRequest()
				target, err := targetFactor.Factorize(ctx, targetResourceReq.TargetId)
				if err != nil {
					return fmt.Errorf("failed to factorize target: %w", err)
				}
				_, err = rh.cli.SendTarget(ctx, &pb.SendTargetRequest{
					RequestId: res.RequestId,
					TargetId:  targetResourceReq.TargetId,
					Target:    target.GetTarget(),
				})
				if err != nil {
					return fmt.Errorf("failed to send target: %w", err)
				}
				log.Info(ctx, "Sent target: %v",
					logger.Value("target_id", targetResourceReq.TargetId))
			case pb.RequestType_REQUEST_TYPE_UNSPECIFIED:
				return fmt.Errorf("request type is unspecified")
			default:
				return fmt.Errorf("unknown request type: %v", res.RequestType)
			}
		}
	}
}
