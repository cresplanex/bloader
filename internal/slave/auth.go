package slave

import (
	"context"
	"fmt"

	"github.com/cresplanex/bloader/internal/auth"
	"github.com/cresplanex/bloader/internal/runner"
	"github.com/cresplanex/bloader/internal/slave/slcontainer"
)

// AuthenticatorFactor represents the slave authenticator factor
type AuthenticatorFactor struct {
	auth                          *slcontainer.Auth
	connectionID                  string
	receiveChanelRequestContainer *slcontainer.ReceiveChanelRequestContainer
	mapper                        *slcontainer.RequestConnectionMapper
}

// Factorize returns the factorized authenticator
func (s *AuthenticatorFactor) Factorize(
	ctx context.Context,
	authID string,
	isDefault bool,
) (auth.SetAuthor, error) {
	if isDefault {
		if s.auth.DefaultAuthenticator == "" {
			term := s.receiveChanelRequestContainer.SendAuthResourceRequests(
				ctx,
				s.connectionID,
				s.mapper,
				slcontainer.AuthResourceRequest{
					IsDefault: true,
				},
			)
			if term == nil {
				return nil, fmt.Errorf("failed to send auth resource request")
			}
			select {
			case <-ctx.Done():
				return nil, nil
			case <-term:
			}
		}
		authID = s.auth.DefaultAuthenticator
	}

	a, ok := s.auth.Find(authID)
	if ok {
		return a, nil
	}

	term := s.receiveChanelRequestContainer.SendAuthResourceRequests(
		ctx,
		s.connectionID,
		s.mapper,
		slcontainer.AuthResourceRequest{
			AuthID: authID,
		},
	)
	if term == nil {
		return nil, fmt.Errorf("failed to send auth resource request")
	}
	select {
	case <-ctx.Done():
		return nil, nil
	case <-term:
	}

	a, ok = s.auth.Find(authID)
	if !ok {
		return nil, fmt.Errorf("auth not found: %s", authID)
	}

	return a, nil
}

// IsDefault returns if the authenticator is the default authenticator
func (s *AuthenticatorFactor) IsDefault(authID string) bool {
	return s.auth.DefaultAuthenticator == authID
}

var _ runner.AuthenticatorFactor = &AuthenticatorFactor{}
