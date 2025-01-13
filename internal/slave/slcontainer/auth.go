package slcontainer

import (
	"sync"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/auth"
)

// Auth is the struct to store auth information
type Auth struct {
	mu                   *sync.RWMutex
	DefaultAuthenticator string
	Container            map[string]*auth.SetAuthor
}

// NewAuth creates a new auth container
func NewAuth() *Auth {
	return &Auth{
		mu:        &sync.RWMutex{},
		Container: make(map[string]*auth.SetAuthor),
	}
}

// Exists checks if the auth exists
func (a Auth) Exists(id string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	_, ok := a.Container[id]
	return ok
}

// Add adds a new auth to the container
func (a *Auth) Add(id string, auth auth.SetAuthor) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Container[id] = &auth
}

// Remove removes an auth from the container
func (a *Auth) Remove(id string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	delete(a.Container, id)
}

// AddFromProto adds a new auth from the proto to the container
func (a Auth) AddFromProto(id string, pbAuth *pb.Auth) error {
	switch pbAuth.Type {
	case pb.AuthType_AUTH_TYPE_OAUTH2:
		auth := &auth.OAuthToken{
			TokenType:   pbAuth.GetOauth2().TokenType,
			AccessToken: pbAuth.GetOauth2().AccessToken,
		}
		a.Add(id, auth)
	case pb.AuthType_AUTH_TYPE_BASIC:
		auth := &auth.BasicAuthenticator{
			Username: pbAuth.GetBasic().Username,
			Password: pbAuth.GetBasic().Password,
		}
		a.Add(id, auth)
	case pb.AuthType_AUTH_TYPE_PRIVATE_KEY:
		// TODO: Implement private key auth
	case pb.AuthType_AUTH_TYPE_API_KEY:
		auth := &auth.APIKeyAuthenticator{
			APIKey:     pbAuth.GetApiKey().ApiKey,
			HeaderName: pbAuth.GetApiKey().HeaderName,
		}
		a.Add(id, auth)
	case pb.AuthType_AUTH_TYPE_JWT:
		auth := &auth.JWTAuthenticator{}
		a.Add(id, auth)
	case pb.AuthType_AUTH_TYPE_UNSPECIFIED:
		return ErrInvalidAuthType
	default:
		return ErrInvalidAuthType
	}
	return nil
}

// Find finds an auth from the container
func (a Auth) Find(id string) (auth.SetAuthor, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	auth, ok := a.Container[id]
	if !ok {
		return nil, false
	}
	return *auth, ok
}

// SetDefault sets the default auth
func (a *Auth) SetDefault(id string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.DefaultAuthenticator = id
}

// GetDefault gets the default auth
func (a Auth) GetDefault() auth.SetAuthor {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return *a.Container[a.DefaultAuthenticator]
}

// AuthResourceRequest is the struct to store auth resource request
type AuthResourceRequest struct {
	IsDefault bool
	AuthID    string
}
