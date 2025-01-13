package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "buf.build/gen/go/cresplanex/bloader/protocolbuffers/go/cresplanex/bloader/v1"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/store"
	"github.com/cresplanex/bloader/internal/utils"
)

// OAuthAuthenticator is an Authenticator that uses OAuth
type OAuthAuthenticator struct {
	grantType            oauthGrantType
	oauthConf            oauth2.Config
	clientcredentialConf clientcredentials.Config
	authCodeOptions      []oauth2.AuthCodeOption
	username             string
	password             string
	redirectPort         int
	credentialConf       config.ValidAuthCredentialConfig
	authToken            *OAuthToken
}

const (
	oauthRedirectHost  = "localhost"
	oauthRedirectPath  = "/auth/callback"
	oauthLoginWaitTime = 60 * time.Second
)

type oauthGrantType int

const (
	_ oauthGrantType = iota
	// AuthOAuth2GrantTypeAuthorizationCode is the OAuth2 grant type Authorization Code
	AuthOAuth2GrantTypeAuthorizationCode
	// AuthOAuth2GrantTypeClientCredentials is the OAuth2 grant type Client Credentials
	AuthOAuth2GrantTypeClientCredentials
	// AuthOAuth2GrantTypePassword is the OAuth2 grant type Password
	AuthOAuth2GrantTypePassword
)

// NewOAuthAuthenticator creates a new OAuthAuthenticator
func NewOAuthAuthenticator(
	str store.Store,
	redirectPort int,
	conf config.ValidAuthOAuth2Config,
) (Authenticator, error) {
	authenticator := &OAuthAuthenticator{}
	switch conf.GrantType {
	case config.AuthOAuth2GrantTypeAuthorizationCode:
		authenticator.grantType = AuthOAuth2GrantTypeAuthorizationCode
		authenticator.oauthConf = oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			Scopes:       conf.Scope,
			RedirectURL:  fmt.Sprintf("http://%s:%d%s", oauthRedirectHost, redirectPort, oauthRedirectPath),
			Endpoint: oauth2.Endpoint{
				AuthURL:  conf.AuthURL,
				TokenURL: conf.TokenURL,
			},
		}
		switch conf.AccessType {
		case config.AuthOAuth2AccessTypeOnline:
			authenticator.authCodeOptions = append(authenticator.authCodeOptions, oauth2.AccessTypeOnline)
		case config.AuthOAuth2AccessTypeOffline:
			authenticator.authCodeOptions = append(authenticator.authCodeOptions, oauth2.AccessTypeOffline)
		}
	case config.AuthOAuth2GrantTypeClientCredentials:
		authenticator.grantType = AuthOAuth2GrantTypeClientCredentials
		authenticator.clientcredentialConf = clientcredentials.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			TokenURL:     conf.TokenURL,
			Scopes:       conf.Scope,
		}
	case config.AuthOAuth2GrantTypePassword:
		authenticator.grantType = AuthOAuth2GrantTypePassword
		authenticator.oauthConf = oauth2.Config{
			ClientID:     conf.ClientID,
			ClientSecret: conf.ClientSecret,
			RedirectURL:  fmt.Sprintf("http://%s:%d%s", oauthRedirectHost, redirectPort, oauthRedirectPath),
			Scopes:       conf.Scope,
			Endpoint: oauth2.Endpoint{
				TokenURL: conf.TokenURL,
			},
		}
		authenticator.username = conf.Username
		authenticator.password = conf.Password
	}
	authenticator.redirectPort = redirectPort
	authenticator.credentialConf = conf.Credential

	authenticator.authToken = &OAuthToken{}
	if authToken, err := credentialGet(str, conf.Credential); err == nil {
		authenticator.authToken = authToken
	}

	return authenticator, nil
}

// OAuthToken is the OAuth token
type OAuthToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	Expiry       time.Time `json:"expiry"`
}

// SetOnRequest sets the authentication information on the request
func (t *OAuthToken) SetOnRequest(_ context.Context, r *http.Request) {
	r.Header.Set("Authorization", t.TokenType+" "+t.AccessToken)
}

// GetAuthValue returns the authentication value
func (t *OAuthToken) GetAuthValue() *pb.Auth {
	return &pb.Auth{
		Type: pb.AuthType_AUTH_TYPE_OAUTH2,
		Auth: &pb.Auth_Oauth2{
			Oauth2: &pb.AuthOAuth2{
				AccessToken: t.AccessToken,
				TokenType:   t.TokenType,
			},
		},
	}
}

// isExpired checks if the token is expired
func (t *OAuthToken) isExpired() bool {
	return t.Expiry.Before(time.Now())
}

func (t *OAuthToken) refresh(
	ctx context.Context,
	str store.Store,
	credentialConf config.ValidAuthCredentialConfig,
	oauthConf oauth2.Config,
) error {
	tokenSource := oauthConf.TokenSource(ctx, &oauth2.Token{
		RefreshToken: t.RefreshToken,
	})
	newToken, err := tokenSource.Token()
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	return credentialSet(t, newToken, str, credentialConf)
}

// Authenticate authenticates the user
func (a *OAuthAuthenticator) Authenticate(ctx context.Context, str store.Store) error {
	switch a.grantType {
	case AuthOAuth2GrantTypeAuthorizationCode:
		var ok bool
		var err error
		var state string
		if state, err = utils.GenerateRandomString(32); err != nil {
			return fmt.Errorf("failed to generate state: %w", err)
		}
		authURL := a.oauthConf.AuthCodeURL(state, a.authCodeOptions...)
		fmt.Println("Please open the following URL in your browser to authenticate:")
		fmt.Printf("Authentication URL:\n+-----------------------------------------------------------+\n\n")
		fmt.Println(authURL)
		fmt.Printf("\n+-----------------------------------------------------------+\n\n")
		forceTimeout := make(chan bool, 1)
		server := &http.Server{
			Addr:              fmt.Sprintf(":%d", a.redirectPort),
			ReadHeaderTimeout: 10 * time.Second,
		}
		http.HandleFunc(oauthRedirectPath, handlerCallbackFactory(ctx, a.oauthConf, func(token *oauth2.Token) {
			fmt.Println("Received token from OAuth server.")
			if err := credentialSet(a.authToken, token, str, a.credentialConf); err != nil {
				fmt.Printf("Failed to save token: %v", err)
				return
			}
			ok = true
		}, forceTimeout, state))

		go func() {
			fmt.Println("Waiting for authentication Callback...", fmt.Sprintf(":%d%s", a.redirectPort, oauthRedirectPath))
			fmt.Println()
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Failed to start server: %v", err)
			}
		}()

		timer := time.NewTimer(oauthLoginWaitTime)

		select {
		case <-forceTimeout:
			fmt.Println("Authentication successful! Shutting down server...")
		case <-timer.C:
			fmt.Println("Timeout reached. Shutting down server...")
		}

		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}
		if !ok {
			return fmt.Errorf("authentication failed")
		}
		return nil
	case AuthOAuth2GrantTypeClientCredentials:
		token, err := a.clientcredentialConf.Token(ctx)
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
		return credentialSet(a.authToken, token, str, a.credentialConf)
	case AuthOAuth2GrantTypePassword:
		token, err := a.oauthConf.PasswordCredentialsToken(ctx, a.username, a.password)
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
		return credentialSet(a.authToken, token, str, a.credentialConf)
	}
	return nil
}

func credentialSet(
	t *OAuthToken,
	token *oauth2.Token,
	str store.Store,
	credentialConf config.ValidAuthCredentialConfig,
) error {
	authToken := OAuthToken{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		Expiry:       token.Expiry,
		RefreshToken: token.RefreshToken,
	}
	authTokenBytes, err := json.Marshal(authToken)
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}
	if err := str.PutObject(
		credentialConf.Store.BucketID,
		credentialConf.Store.Key,
		authTokenBytes,
	); err != nil {
		return fmt.Errorf("failed to save token to store: %w", err)
	}
	*t = authToken
	return nil
}

func credentialGet(str store.Store, credentialConf config.ValidAuthCredentialConfig) (*OAuthToken, error) {
	authTokenBytes, err := str.GetObject(credentialConf.Store.BucketID, credentialConf.Store.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to get token from store: %w", err)
	}
	if len(authTokenBytes) == 0 {
		return &OAuthToken{}, nil
	}
	authToken := &OAuthToken{}
	if err := json.Unmarshal(authTokenBytes, authToken); err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}
	return authToken, nil
}

// SetOnRequest sets the authentication information on the request
func (a *OAuthAuthenticator) SetOnRequest(ctx context.Context, r *http.Request) {
	a.authToken.SetOnRequest(ctx, r)
}

// GetAuthValue returns the authentication value
func (a *OAuthAuthenticator) GetAuthValue() *pb.Auth {
	return &pb.Auth{
		Type: pb.AuthType_AUTH_TYPE_OAUTH2,
		Auth: &pb.Auth_Oauth2{
			Oauth2: &pb.AuthOAuth2{
				AccessToken: a.authToken.AccessToken,
				TokenType:   a.authToken.TokenType,
			},
		},
	}
}

// IsExpired checks if the authentication information is expired
func (a *OAuthAuthenticator) IsExpired(_ context.Context, _ store.Store) bool {
	return a.authToken.isExpired()
}

// Refresh refreshes the authentication information
func (a *OAuthAuthenticator) Refresh(ctx context.Context, str store.Store) error {
	switch a.grantType {
	case AuthOAuth2GrantTypeAuthorizationCode:
		return a.authToken.refresh(ctx, str, a.credentialConf, a.oauthConf)
	case AuthOAuth2GrantTypeClientCredentials:
		token, err := a.clientcredentialConf.Token(ctx)
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
		return credentialSet(a.authToken, token, str, a.credentialConf)
	case AuthOAuth2GrantTypePassword:
		token, err := a.oauthConf.PasswordCredentialsToken(ctx, a.username, a.password)
		if err != nil {
			return fmt.Errorf("failed to get token: %w", err)
		}
		return credentialSet(a.authToken, token, str, a.credentialConf)
	}
	return nil
}

var _ Authenticator = &OAuthAuthenticator{}

func handlerCallbackFactory(
	ctx context.Context,
	oauthConf oauth2.Config,
	setter func(token *oauth2.Token),
	shutdownFlag chan<- bool,
	state string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "Invalid state parameter", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")
		token, err := oauthConf.Exchange(ctx, code)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to exchange token: %v", err), http.StatusInternalServerError)
			return
		}

		setter(token)
		fmt.Fprintf(w, "Authentication successful! You can close this window.")

		shutdownFlag <- true
	}
}

var (
	_ SetAuthor     = &OAuthToken{}
	_ SetAuthor     = &OAuthAuthenticator{}
	_ Authenticator = &OAuthAuthenticator{}
)
