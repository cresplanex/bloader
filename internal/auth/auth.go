// Package auth provides the authentication logic for the application
package auth

import (
	"context"
	"net/http"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/store"
)

// SetAuthor sets the author on the request
type SetAuthor interface {
	// SetOnRequest sets the authentication information on the request
	SetOnRequest(ctx context.Context, r *http.Request)
	// GetAuthValue returns the authentication value
	GetAuthValue() *pb.Auth
}

// Authenticator is an interface for authenticating
type Authenticator interface {
	SetAuthor
	// Authenticate authenticates the user
	Authenticate(ctx context.Context, str store.Store) error
	// IsExpired checks if the authentication information is expired
	IsExpired(ctx context.Context, str store.Store) bool
	// Refresh refreshes the authentication information
	Refresh(ctx context.Context, str store.Store) error
}

// AuthenticatorContainer holds the dependencies for the Authenticator
type AuthenticatorContainer struct {
	DefaultAuthenticator string
	Container            map[string]*Authenticator
}

// NewAuthenticatorContainerFromConfig creates a new AuthenticatorContainer from the configuration
func NewAuthenticatorContainerFromConfig(str store.Store, conf config.ValidConfig) (AuthenticatorContainer, error) {
	ctr := AuthenticatorContainer{
		Container: make(map[string]*Authenticator),
	}

	for _, authConf := range conf.Auth {
		var authenticator Authenticator
		var err error
		switch authConf.Type {
		case config.AuthTypeOAuth2:
			var redirectPort int
			if conf.Server.RedirectPort.Enabled {
				redirectPort = conf.Server.RedirectPort.Port
			} else {
				redirectPort = conf.Server.Port
			}
			authenticator, err = NewOAuthAuthenticator(str, redirectPort, authConf.OAuth2)
			if err != nil {
				return AuthenticatorContainer{}, err
			}
		case config.AuthTypeBasic:
			authenticator, err = NewBasicAuthenticator(authConf.Basic)
			if err != nil {
				return AuthenticatorContainer{}, err
			}
		case config.AuthTypeAPIKey:
			authenticator, err = NewAPIKeyAuthenticator(authConf.APIKey)
			if err != nil {
				return AuthenticatorContainer{}, err
			}
		case config.AuthTypePrivateKey:
			// TODO: Implement PrivateKeyAuthenticator
		case config.AuthTypeJWT:
			authenticator, err = NewJWTAuthenticator(authConf.JWT)
			if err != nil {
				return AuthenticatorContainer{}, err
			}
		}

		ctr.Container[authConf.ID] = &authenticator
		if authConf.Default {
			ctr.DefaultAuthenticator = authConf.ID
		}
	}

	return ctr, nil
}
