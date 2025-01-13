package slcontainer

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/runner"
	"github.com/cresplanex/bloader/internal/utils"
)

// RequestTermCaster is an interface that represents a request term caster
type RequestTermCaster struct {
	mu  *sync.RWMutex
	req map[string]chan<- struct{}
}

// NewRequestTermCaster creates a new request term caster
func NewRequestTermCaster() *RequestTermCaster {
	return &RequestTermCaster{
		mu:  &sync.RWMutex{},
		req: make(map[string]chan<- struct{}),
	}
}

// RegisterRequest registers a new request to the caster
func (r *RequestTermCaster) RegisterRequest(reqID string) <-chan struct{} {
	r.mu.Lock()
	defer r.mu.Unlock()

	ch := make(chan struct{})
	r.req[reqID] = ch
	return ch
}

// Cast casts a term to the request
func (r *RequestTermCaster) Cast(reqID string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if ch, ok := r.req[reqID]; ok {
		close(ch)
	}
}

// ReceiveChanelRequestContainer is a struct that represents a container for the receive chanel requests
type ReceiveChanelRequestContainer struct {
	mu         *sync.RWMutex
	termCaster *RequestTermCaster
	ReqChan    chan *pb.ReceiveChanelConnectResponse
}

// NewReceiveChanelRequestContainer creates a new request container
func NewReceiveChanelRequestContainer() *ReceiveChanelRequestContainer {
	return &ReceiveChanelRequestContainer{
		mu:         &sync.RWMutex{},
		termCaster: NewRequestTermCaster(),
		ReqChan:    make(chan *pb.ReceiveChanelConnectResponse),
	}
}

// SendLoaderResourceRequests sets the loader requests channel
func (r *ReceiveChanelRequestContainer) SendLoaderResourceRequests(
	ctx context.Context,
	connectionID string,
	mapper *RequestConnectionMapper,
	req LoaderResourceRequest,
) <-chan struct{} {
	r.mu.Lock()
	defer r.mu.Unlock()

	requestID := utils.GenerateUniqueID()

	pbReq := &pb.ReceiveChanelConnectResponse{
		RequestId:   requestID,
		RequestType: pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_LOADER,
		Request: &pb.ReceiveChanelConnectResponse_LoaderResourceRequest{
			LoaderResourceRequest: &pb.ReceiveChanelConnectLoaderResourceRequest{
				LoaderId: req.LoaderID,
			},
		},
	}

	select {
	case <-ctx.Done():
		return nil
	case r.ReqChan <- pbReq: // nothing
	}

	mapper.RegisterRequestConnection(requestID, connectionID)

	return r.termCaster.RegisterRequest(requestID)
}

// SendAuthResourceRequests sets the auth requests channel
func (r *ReceiveChanelRequestContainer) SendAuthResourceRequests(
	ctx context.Context,
	connectionID string,
	mapper *RequestConnectionMapper,
	req AuthResourceRequest,
) <-chan struct{} {
	r.mu.Lock()
	defer r.mu.Unlock()

	requestID := utils.GenerateUniqueID()

	pbReq := &pb.ReceiveChanelConnectResponse{
		RequestId:   requestID,
		RequestType: pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_AUTH,
		Request: &pb.ReceiveChanelConnectResponse_AuthResourceRequest{
			AuthResourceRequest: &pb.ReceiveChanelConnectAuthResourceRequest{
				AuthId:    req.AuthID,
				IsDefault: req.IsDefault,
			},
		},
	}

	select {
	case <-ctx.Done():
		return nil
	case r.ReqChan <- pbReq: // nothing
	}

	mapper.RegisterRequestConnection(requestID, connectionID)

	return r.termCaster.RegisterRequest(requestID)
}

// SendStore send store requests
func (r *ReceiveChanelRequestContainer) SendStore(
	ctx context.Context,
	connectionID string,
	mapper *RequestConnectionMapper,
	req StoreDataRequest,
) (<-chan struct{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	requestID := utils.GenerateUniqueID()

	strData := make([]*pb.StoreData, 0, len(req.StoreData))

	for _, storeData := range req.StoreData {
		data := &pb.StoreData{
			BucketId: storeData.BucketID,
			StoreKey: storeData.StoreKey,
			Data:     storeData.Data,
		}
		if storeData.Encryption.Enabled {
			data.Encryption = &pb.Encryption{
				Enabled:   storeData.Encryption.Enabled,
				EncryptId: storeData.Encryption.EncryptID,
			}
		}
		strData = append(strData, data)
	}

	strDataList := &pb.StoreDataList{
		Data: strData,
	}

	b, err := proto.Marshal(strDataList)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal store data list: %w", err)
	}
	uid := utils.GenerateUniqueID()

	for i := 0; i < len(b); i += runner.DefaultChunkSize {
		end := i + runner.DefaultChunkSize
		if end > len(b) {
			end = len(b)
		}
		pbReq := &pb.ReceiveChanelConnectResponse{
			RequestId:   requestID,
			RequestType: pb.RequestType_REQUEST_TYPE_STORE,
			Request: &pb.ReceiveChanelConnectResponse_Store{
				Store: &pb.ReceiveChanelConnectStore{
					Uid:         uid,
					Data:        b[i:end],
					IsLastChunk: end == len(b),
				},
			},
		}
		select {
		case <-ctx.Done():
		case r.ReqChan <- pbReq: // nothing
		}
	}

	mapper.RegisterRequestConnection(requestID, connectionID)

	return r.termCaster.RegisterRequest(requestID), nil
}

// SendStoreResourceRequests sets the store requests channel
func (r *ReceiveChanelRequestContainer) SendStoreResourceRequests(
	ctx context.Context,
	connectionID string,
	mapper *RequestConnectionMapper,
	req StoreResourceRequest,
) (<-chan struct{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	requestID := utils.GenerateUniqueID()

	strData := make([]*pb.StoreImportRequest, 0, len(req.Requests))

	for _, storeData := range req.Requests {
		data := &pb.StoreImportRequest{
			BucketId: storeData.BucketID,
			StoreKey: storeData.StoreKey,
		}
		if storeData.Encryption.Enabled {
			data.Encryption = &pb.Encryption{
				Enabled:   storeData.Encryption.Enabled,
				EncryptId: storeData.Encryption.EncryptID,
			}
		}
		strData = append(strData, data)
	}

	strDataList := &pb.StoreImportRequestList{
		Data: strData,
	}

	b, err := proto.Marshal(strDataList)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal store data list: %w", err)
	}
	uid := utils.GenerateUniqueID()

	for i := 0; i < len(b); i += runner.DefaultChunkSize {
		end := i + runner.DefaultChunkSize
		if end > len(b) {
			end = len(b)
		}
		pbReq := &pb.ReceiveChanelConnectResponse{
			RequestId:   requestID,
			RequestType: pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_STORE,
			Request: &pb.ReceiveChanelConnectResponse_StoreResourceRequest{
				StoreResourceRequest: &pb.ReceiveChanelConnectStoreResourceRequest{
					Uid:         uid,
					Data:        b[i:end],
					IsLastChunk: end == len(b),
				},
			},
		}
		select {
		case <-ctx.Done():
		case r.ReqChan <- pbReq: // nothing
		}
	}

	mapper.RegisterRequestConnection(requestID, connectionID)

	return r.termCaster.RegisterRequest(requestID), nil
}

// SendTargetResourceRequests sets the target requests channel
func (r *ReceiveChanelRequestContainer) SendTargetResourceRequests(
	ctx context.Context,
	connectionID string,
	mapper *RequestConnectionMapper,
	req TargetResourceRequest,
) <-chan struct{} {
	r.mu.Lock()
	defer r.mu.Unlock()

	requestID := utils.GenerateUniqueID()

	pbReq := &pb.ReceiveChanelConnectResponse{
		RequestId:   requestID,
		RequestType: pb.RequestType_REQUEST_TYPE_REQUEST_RESOURCE_TARGET,
		Request: &pb.ReceiveChanelConnectResponse_TargetResourceRequest{
			TargetResourceRequest: &pb.ReceiveChanelConnectTargetResourceRequest{
				TargetId: req.TargetID,
			},
		},
	}

	select {
	case <-ctx.Done():
		return nil
	case r.ReqChan <- pbReq: // nothing
	}

	mapper.RegisterRequestConnection(requestID, connectionID)

	return r.termCaster.RegisterRequest(requestID)
}

// Cast casts a term to the request
func (r *ReceiveChanelRequestContainer) Cast(reqID string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.termCaster.Cast(reqID)
}
