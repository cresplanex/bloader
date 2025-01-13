package slave

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cresplanex/bloader/internal/encrypt"
)

// UnaryServerEncryptInterceptor is a server-side interceptor that encrypts the request and decrypts the response.
func UnaryServerEncryptInterceptor(encrypter encrypt.Encrypter) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if encryptedReq, ok := req.(string); ok {
			plainReq, err := encrypter.Decrypt(encryptedReq)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to decrypt request: %v", err)
			}
			req = plainReq
		}

		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		if plainResp, ok := resp.([]byte); ok {
			encryptedResp, err := encrypter.Encrypt(plainResp)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to encrypt response: %v", err)
			}
			resp = encryptedResp
		}

		return resp, nil
	}
}

// StreamServerInterceptor is a server-side interceptor that encrypts the request and decrypts the response.
func StreamServerInterceptor(encrypter encrypt.Encrypter) grpc.StreamServerInterceptor {
	return func(
		srv any,
		ss grpc.ServerStream,
		_ *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		wrappedStream := &wrappedServerStream{
			ServerStream: ss,
			encrypter:    encrypter,
		}
		return handler(srv, wrappedStream)
	}
}

type wrappedServerStream struct {
	grpc.ServerStream
	encrypter encrypt.Encrypter
}

// RecvMsg implements the grpc.ServerStream interface.
func (w *wrappedServerStream) RecvMsg(m any) error {
	var encryptedMsg []byte
	if err := w.ServerStream.RecvMsg(&encryptedMsg); err != nil {
		return fmt.Errorf("failed to receive message: %w", err)
	}
	descryptedMsg, err := w.encrypter.Decrypt(string(encryptedMsg))
	if err != nil {
		return status.Errorf(codes.Internal, "failed to decrypt request: %v", err)
	}
	if byteMsg, ok := m.(*[]byte); ok {
		*byteMsg = descryptedMsg
		return nil
	}

	return fmt.Errorf("failed to cast message to byte")
}

// SendMsg implements the grpc.ServerStream interface.
func (w *wrappedServerStream) SendMsg(m any) error {
	if plainMsg, ok := m.([]byte); ok {
		encryptedMsg, err := w.encrypter.Encrypt(plainMsg)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to encrypt response: %v", err)
		}
		m = encryptedMsg
	}
	if err := w.ServerStream.SendMsg(m); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
