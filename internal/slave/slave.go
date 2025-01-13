// Package slave provides the slave node of the bloader.
package slave

import (
	"fmt"
	"net"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/cresplanex/bloader/internal/container"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/runner"
)

// Run runs the slave node
func Run(ctr *container.Container) error {
	var grpcServerOptions []grpc.ServerOption
	if ctr.Config.SlaveSetting.Certificate.Enabled {
		creds, err := credentials.NewServerTLSFromFile(
			ctr.Config.SlaveSetting.Certificate.SlaveCert,
			ctr.Config.SlaveSetting.Certificate.SlaveKey,
		)
		if err != nil {
			return fmt.Errorf("failed to load certificate: %w", err)
		}
		grpcServerOptions = append(grpcServerOptions, grpc.Creds(creds))
	}

	if ctr.Config.SlaveSetting.Encrypt.Enabled {
		encrypter, ok := ctr.EncypterContainer[ctr.Config.SlaveSetting.Encrypt.EncryptID]
		if !ok {
			return fmt.Errorf("encrypter not found: %s", ctr.Config.SlaveSetting.Encrypt.EncryptID)
		}
		grpcServerOptions = append(
			grpcServerOptions,
			grpc.UnaryInterceptor(UnaryServerEncryptInterceptor(encrypter)),
			grpc.StreamInterceptor(StreamServerInterceptor(encrypter)),
		)
	}

	grpcServer := grpc.NewServer(grpcServerOptions...)

	slCtr := runner.NewConnectionContainer()
	defer slCtr.AllDisconnect(ctr.Ctx)

	pb.RegisterBloaderSlaveServiceServer(grpcServer, NewServer(ctr, slCtr))
	lister, err := net.Listen("tcp", fmt.Sprintf(":%d", ctr.Config.SlaveSetting.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	ctr.Logger.Info(ctr.Ctx, "Starting the worker node",
		logger.Value("port", ctr.Config.SlaveSetting.Port))

	go func() {
		<-ctr.Ctx.Done()
		ctr.Logger.Info(ctr.Ctx, "Shutting down the worker node")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lister); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
