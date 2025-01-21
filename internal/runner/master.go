package runner

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/encrypt"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/master"
)

// ReceiveTermData is a struct that holds the receive term information.
type ReceiveTermData struct {
	Type ReceiveTermType
	Err  error
}

// ReceiveTermType represents the valid ReceiveTermType runner
type ReceiveTermType string

// ReceiveTermType constants
const (
	// ReceiveTermTypeEOF represents the EOF
	ReceiveTermTypeReceiveTermTypeEOF ReceiveTermType = "EOF"
	// ReceiveTermTypeResponseReceiveError represents the ResponseReceiveError
	ReceiveTermTypeReceiveTermTypeResponseReceiveError ReceiveTermType = "ResponseReceiveError"
	// ReceiveTermTypeStreamContextDone represents the StreamContextDone
	ReceiveTermTypeReceiveTermTypeStreamContextDone ReceiveTermType = "StreamContextDone"
	// ReceiveTermTypeDisconnected represents the Disconnected
	ReceiveTermTypeReceiveTermTypeDisconnected ReceiveTermType = "Disconnected"
)

// ConnectionMapData is a struct that holds the connection information.
type ConnectionMapData struct {
	ConnectionID    string
	conn            *grpc.ClientConn
	Cli             pb.BloaderSlaveServiceClient
	ReqChan         <-chan *pb.ReceiveChanelConnectResponse
	termChan        chan<- struct{}
	ReceiveTermChan <-chan ReceiveTermData
}

// ConnectionContainer is a struct that holds the connection information.
type ConnectionContainer struct {
	mu     *sync.RWMutex
	conMap map[string]*ConnectionMapData // Key: slaveID
}

// NewConnectionContainer creates a new ConnectMap.
func NewConnectionContainer() *ConnectionContainer {
	return &ConnectionContainer{
		mu:     &sync.RWMutex{},
		conMap: make(map[string]*ConnectionMapData),
	}
}

// Find returns the connection information.
func (c *ConnectionContainer) Find(slaveID string) (*ConnectionMapData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	conn, ok := c.conMap[slaveID]
	if !ok {
		return nil, false
	}

	return conn, true
}

// isErrConnectionReset checks if the error is connection reset.
func isErrConnectionReset(err error) bool {
	// 	Connection reset with read error.
	// Can retry because of power equality.
	if strings.Contains(err.Error(), "read: connection reset") {
		return true
	}

	if strings.Contains(err.Error(), "use of closed network connection") ||
		strings.Contains(err.Error(), "connection reset") ||
		strings.Contains(err.Error(), "broken pipe") {
		return true
	}

	return false
}

// Connect adds a connection to the map.
func (c *ConnectionContainer) Connect(
	ctx context.Context,
	log logger.Logger,
	env string,
	encryptCtr encrypt.Container,
	conInfo ValidSlaveConnect,
	eventCaster EventCaster,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := eventCaster.CastEvent(ctx, SlaveConnectRunnerEventConnecting); err != nil {
		return fmt.Errorf("failed to cast event: %w", err)
	}
	defer func() {
		if err := eventCaster.CastEvent(ctx, SlaveConnectRunnerEventConnected); err != nil {
			log.Error(ctx, "failed to cast event: %v", logger.Value("error", err))
		}
	}()

	for _, slave := range conInfo.Slaves {
		if _, ok := c.conMap[slave.ID]; ok {
			return fmt.Errorf("connection already exists: %s", slave.ID)
		}
		grpcDialOptions := []grpc.DialOption{}
		if slave.Certificate.Enabled {
			b, err := os.ReadFile(slave.Certificate.CACert)
			if err != nil {
				return fmt.Errorf("credentials: failed to read CA certificate: %w", err)
			}
			cp := x509.NewCertPool()
			if !cp.AppendCertsFromPEM(b) {
				return fmt.Errorf("credentials: failed to append certificates")
			}
			creds := credentials.NewTLS(&tls.Config{
				ServerName: slave.Certificate.ServerNameOverride,
				//nolint:gosec
				InsecureSkipVerify: slave.Certificate.InsecureSkipVerify,
				RootCAs:            cp,
			})
			grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(creds))
		} else {
			grpcDialOptions = append(grpcDialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		if slave.Encrypt.Enabled {
			encrypter, ok := encryptCtr[slave.Encrypt.EncryptID]
			if !ok {
				return fmt.Errorf("encrypter not found: %s", slave.Encrypt.EncryptID)
			}
			grpcDialOptions = append(
				grpcDialOptions,
				grpc.WithUnaryInterceptor(master.UnaryClientEncryptInterceptor(encrypter)),
				grpc.WithStreamInterceptor(master.StreamClientInterceptor(encrypter)),
			)
		}

		// retryPolicy := `{
		// 	"loadBalancingConfig": [ { "round_robin": {} } ],

		//     "methodConfig": [{
		//         "name": [{"service": "cresplanex.bloader.v1.BloaderSlaveService"}],

		//         "retryPolicy": {
		//             "MaxAttempts": 5,
		//             "InitialBackoff": ".05s",
		//             "MaxBackoff": "1s",
		//             "BackoffMultiplier": 2.0,
		//             "RetryableStatusCodes": [ "UNAVAILABLE" ]
		//         }
		//     }]
		// }`

		// grpcDialOptions = append(
		// 	grpcDialOptions,
		// 	grpc.WithDefaultServiceConfig(retryPolicy),
		// 	// grpc.WithKeepaliveParams(
		// 	// 	keepalive.ClientParameters{
		// 	// 		Time:                10 * time.Second, // TODO: Set the keepalive parameters from config
		// 	// 		Timeout:             30 * time.Second, // TODO: Set the keepalive parameters from config
		// 	// 		PermitWithoutStream: true,
		// 	// 	},
		// 	// ),
		// )

		conn, err := grpc.NewClient(slave.URI, grpcDialOptions...)
		if err != nil {
			return fmt.Errorf("failed to connect to slave: %w", err)
		}

		cli := pb.NewBloaderSlaveServiceClient(conn)

		conReq := &pb.ConnectRequest{
			Environment: env,
		}

		res, err := cli.Connect(ctx, conReq)
		if err != nil {
			return fmt.Errorf("failed to connect to slave: %w", err)
		}

		conID := res.ConnectionId

		receiveStream, err := cli.ReceiveChanelConnect(
			ctx,
			&pb.ReceiveChanelConnectRequest{
				ConnectionId: conID,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to receive channel connect: %w", err)
		}

		reqChan := make(chan *pb.ReceiveChanelConnectResponse)
		receiveTermChan := make(chan ReceiveTermData)
		termChan := make(chan struct{})

		ctx, cancel := context.WithCancel(ctx)

		go func() {
			defer cancel()
			defer close(reqChan)
			defer close(receiveTermChan)

			// TODO: Set the retry policy from config
			maxAttempts := 5
			retryInterval := 2 * time.Second
			attempts := 0

			for {
				res, err := receiveStream.Recv()
				if errors.Is(err, context.Canceled) {
					return
				}
				if errors.Is(err, io.EOF) {
					log.Info(ctx, "receiveChan EOF")
					select {
					case <-ctx.Done():
						log.Info(ctx, "context done")
						return
					case receiveTermChan <- ReceiveTermData{
						Type: ReceiveTermTypeReceiveTermTypeEOF,
						Err:  nil,
					}:
						log.Info(ctx, "receiveChan EOF")
					}
					return
				}
				if err != nil {
					st, ok := status.FromError(err)
					if ok && st.Code() == codes.Canceled {
						log.Info(ctx, "context canceled rpc error")
						return
					}
					if errors.Is(err, context.Canceled) {
						log.Info(ctx, "context done")
						return
					}

					log.Error(ctx, "failed to receive channel connect",
						logger.Value("error", err), logger.Value("slaveID", slave.ID))

					// Retry the connection when err is reset by peer
					if isErrConnectionReset(err) && attempts < maxAttempts {
						attempts++
						log.Warn(ctx, "retrying connection",
							logger.Value("attempts", attempts),
							logger.Value("maxAttempts", maxAttempts),
							logger.Value("retryInterval", retryInterval),
							logger.Value("slaveID", slave.ID),
							logger.Value("error", err),
							logger.Value("res", res),
						)

						time.Sleep(retryInterval)

						// Do I need the following?
						// receiveStream, err = cli.ReceiveChanelConnect(
						// 	ctx,
						// 	&pb.ReceiveChanelConnectRequest{
						// 		ConnectionId: conID,
						// 	},
						// )
						// if err != nil {
						// 	log.Error(ctx, "failed to retry connection: %v", logger.Value("error", err))
						// 	continue
						// }
						continue
					}

					if err := receiveStream.CloseSend(); err != nil {
						log.Error(ctx, "failed to close receiveChan: %v", logger.Value("error", err))
					}
					select {
					case <-ctx.Done():
						return
					case receiveTermChan <- ReceiveTermData{
						Type: ReceiveTermTypeReceiveTermTypeResponseReceiveError,
						Err:  err,
					}:
					}

					return
				}

				attempts = 0
				select {
				case <-ctx.Done():
					log.Info(ctx, "context done")
					if err := receiveStream.CloseSend(); err != nil {
						log.Error(ctx, "failed to close receiveChan: %v", logger.Value("error", err))
					}
					return
				case <-receiveStream.Context().Done():
					log.Info(ctx, "receiveChan context done")
					select {
					case <-ctx.Done():
						log.Info(ctx, "context done")
						return
					case receiveTermChan <- ReceiveTermData{
						Type: ReceiveTermTypeReceiveTermTypeStreamContextDone,
						Err:  nil,
					}:
					}
					return
				case <-termChan:
					if err := receiveStream.CloseSend(); err != nil {
						log.Error(ctx, "failed to close receiveChan: %v", logger.Value("error", err))
					}
					select {
					case <-ctx.Done():
						return
					case receiveTermChan <- ReceiveTermData{
						Type: ReceiveTermTypeReceiveTermTypeDisconnected,
						Err:  nil,
					}:
					}
					return
				case reqChan <- res:
				}
			}
		}()

		c.conMap[slave.ID] = &ConnectionMapData{
			ConnectionID:    res.ConnectionId,
			conn:            conn,
			Cli:             cli,
			ReqChan:         reqChan,
			termChan:        termChan,
			ReceiveTermChan: receiveTermChan,
		}
	}

	return nil
}

// disconnect removes a connection from the map.
func (c *ConnectionContainer) disconnect(slaveID string) error {
	conn, ok := c.conMap[slaveID]
	if !ok {
		return fmt.Errorf("connection not found: %s", slaveID)
	}
	close(conn.termChan)
	disReq := &pb.DisconnectRequest{
		ConnectionId: conn.ConnectionID,
	}

	_, err := conn.Cli.Disconnect(context.Background(), disReq)
	if err != nil {
		return fmt.Errorf("failed to disconnect from slave: %w", err)
	}
	if err := conn.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	delete(c.conMap, slaveID)

	return nil
}

// Disconnect removes a connection from the map.
func (c *ConnectionContainer) Disconnect(_ context.Context, slaveIDs []string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, slaveID := range slaveIDs {
		if err := c.disconnect(slaveID); err != nil {
			return fmt.Errorf("failed to disconnect from slave: %w", err)
		}
	}

	return nil
}

// AllDisconnect removes all connections from the map.
func (c *ConnectionContainer) AllDisconnect(_ context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for slaveID := range c.conMap {
		if err := c.disconnect(slaveID); err != nil {
			return fmt.Errorf("failed to disconnect from slave: %w", err)
		}
	}

	return nil
}
