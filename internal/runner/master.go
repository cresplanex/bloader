package runner

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	rpc "buf.build/gen/go/cresplanex/bloader/grpc/go/cresplanex/bloader/v1/bloaderv1grpc"
	pb "buf.build/gen/go/cresplanex/bloader/protocolbuffers/go/cresplanex/bloader/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cresplanex/bloader/internal/encrypt"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/master"
)

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
	Cli             rpc.BloaderSlaveServiceClient
	ReqChan         <-chan *pb.ReceiveChanelConnectResponse
	termChan        chan<- struct{}
	ReceiveTermChan <-chan ReceiveTermType
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

		conn, err := grpc.NewClient(slave.URI, grpcDialOptions...)
		if err != nil {
			return fmt.Errorf("failed to connect to slave: %w", err)
		}

		cli := rpc.NewBloaderSlaveServiceClient(conn)

		conReq := &pb.ConnectRequest{
			Environment: env,
		}

		res, err := cli.Connect(ctx, conReq)
		if err != nil {
			return fmt.Errorf("failed to connect to slave: %w", err)
		}

		receiveStream, err := cli.ReceiveChanelConnect(
			ctx,
			&pb.ReceiveChanelConnectRequest{
				ConnectionId: res.ConnectionId,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to receive channel connect: %w", err)
		}

		reqChan := make(chan *pb.ReceiveChanelConnectResponse)
		receiveTermChan := make(chan ReceiveTermType)
		termChan := make(chan struct{})

		ctx, cancel := context.WithCancel(ctx)

		go func() {
			defer cancel()
			defer close(reqChan)
			defer close(receiveTermChan)

			for {
				res, err := receiveStream.Recv()
				if errors.Is(err, io.EOF) {
					log.Info(ctx, "receiveChan EOF")
					select {
					case <-ctx.Done():
						log.Info(ctx, "context done")
						return
					case receiveTermChan <- ReceiveTermTypeReceiveTermTypeEOF:
						log.Info(ctx, "receiveChan EOF")
					}

					return
				}
				if err != nil {
					log.Error(ctx, "failed to receive channel connect: %v",
						logger.Value("error", err), logger.Value("slaveID", slave.ID))
					if err := receiveStream.CloseSend(); err != nil {
						log.Error(ctx, "failed to close receiveChan: %v", logger.Value("error", err))
					}
					select {
					case <-ctx.Done():
						log.Info(ctx, "context done")
						return
					case receiveTermChan <- ReceiveTermTypeReceiveTermTypeResponseReceiveError:
						log.Info(ctx, "receiveChan response receive error")
					}

					return
				}
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
					case receiveTermChan <- ReceiveTermTypeReceiveTermTypeStreamContextDone:
						log.Info(ctx, "receiveChan context done")
					}
					return
				case <-termChan:
					log.Info(ctx, "termChan")
					if err := receiveStream.CloseSend(); err != nil {
						log.Error(ctx, "failed to close receiveChan: %v", logger.Value("error", err))
					}
					select {
					case <-ctx.Done():
						log.Info(ctx, "context done")
						return
					case receiveTermChan <- ReceiveTermTypeReceiveTermTypeDisconnected:
						log.Info(ctx, "receiveChan disconnected")
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
