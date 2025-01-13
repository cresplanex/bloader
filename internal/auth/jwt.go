package auth

import (
	"context"
	"net/http"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/store"
)

// JWTAuthenticator is an Authenticator that uses Private Key
type JWTAuthenticator struct {
	credentialConf config.ValidAuthCredentialConfig
}

// NewJWTAuthenticator creates a new JWTAuthenticator
func NewJWTAuthenticator(conf config.ValidAuthJWTConfig) (Authenticator, error) {
	return &JWTAuthenticator{
		credentialConf: conf.Credential,
	}, nil
}

// Authenticate authenticates the user
func (a *JWTAuthenticator) Authenticate(_ context.Context, _ store.Store) error {
	return nil
}

// SetOnRequest sets the authentication information on the request
func (a *JWTAuthenticator) SetOnRequest(_ context.Context, _ *http.Request) {
}

// GetAuthValue returns the authentication value
func (a *JWTAuthenticator) GetAuthValue() *pb.Auth {
	return &pb.Auth{
		Type: pb.AuthType_AUTH_TYPE_JWT,
		Auth: &pb.Auth_Jwt{
			Jwt: &pb.AuthJwt{
				Jwt: "", // TODO: Implement
			},
		},
	}
}

// IsExpired checks if the authentication information is expired
func (a *JWTAuthenticator) IsExpired(_ context.Context, _ store.Store) bool {
	return false
}

// Refresh refreshes the authentication information
func (a *JWTAuthenticator) Refresh(_ context.Context, _ store.Store) error {
	return nil
}

var (
	_ Authenticator = &JWTAuthenticator{}
	_ SetAuthor     = &JWTAuthenticator{}
)
