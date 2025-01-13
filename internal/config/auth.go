package config

import "fmt"

// AuthRespectiveConfig is the configuration for the respective auth.
type AuthRespectiveConfig struct {
	ID         *string               `mapstructure:"id"`
	Default    bool                  `mapstructure:"default"`
	Type       *string               `mapstructure:"type"`
	OAuth2     *AuthOAuth2Config     `mapstructure:"oauth2"`
	APIKey     *AuthAPIKeyConfig     `mapstructure:"api_key"`
	Basic      *AuthBasicConfig      `mapstructure:"basic"`
	PrivateKey *AuthPrivateKeyConfig `mapstructure:"private_key"`
	JWT        *AuthJWTConfig        `mapstructure:"jwt"`
}

// ValidAuthRespectiveConfig represents the valid auth configuration
type ValidAuthRespectiveConfig struct {
	ID         string
	Default    bool
	Type       AuthType
	OAuth2     ValidAuthOAuth2Config
	APIKey     ValidAuthAPIKeyConfig
	Basic      ValidAuthBasicConfig
	PrivateKey ValidAuthPrivateKeyConfig
	JWT        ValidAuthJWTConfig
}

// AuthType is the type of the auth.
type AuthType string

const (
	// AuthTypeOAuth2 is the type of the oauth2.
	AuthTypeOAuth2 AuthType = "oauth2"
	// AuthTypeAPIKey is the type of the api key.
	AuthTypeAPIKey AuthType = "apiKey"
	// AuthTypeBasic is the type of the basic auth.
	AuthTypeBasic AuthType = "basic"
	// AuthTypePrivateKey is the type of the private key.
	AuthTypePrivateKey AuthType = "privateKey"
	// AuthTypeJWT is the type of the jwt.
	AuthTypeJWT AuthType = "jwt"
)

// AuthCredentialConfig is the configuration for the credential.
type AuthCredentialConfig struct {
	Store *StoreSpecifyConfig `mapstructure:"store"`
}

// ValidAuthCredentialConfig represents the valid auth credential configuration
type ValidAuthCredentialConfig struct {
	Store ValidStoreSpecifyConfig
}

// Validate validates the credential configuration.
func (c AuthCredentialConfig) Validate() (ValidAuthCredentialConfig, error) {
	var valid ValidAuthCredentialConfig
	var err error
	if c.Store == nil {
		return ValidAuthCredentialConfig{}, ErrAuthCredentialStoreRequired
	}
	valid.Store, err = c.Store.Validate()
	if err != nil {
		return ValidAuthCredentialConfig{}, fmt.Errorf("store: %w", err)
	}
	return valid, nil
}

// AuthOAuth2GrantType is the grant type of the oauth2.
type AuthOAuth2GrantType string

const (
	// AuthOAuth2GrantTypeAuthorizationCode is the grant type of the authorization code.
	AuthOAuth2GrantTypeAuthorizationCode AuthOAuth2GrantType = "authorization_code"
	// AuthOAuth2GrantTypeClientCredentials is the grant type of the client credentials.
	AuthOAuth2GrantTypeClientCredentials AuthOAuth2GrantType = "client_credentials"
	// AuthOAuth2GrantTypePassword is the grant type of the password.
	AuthOAuth2GrantTypePassword AuthOAuth2GrantType = "password"
)

// AuthOAuth2AccessType is the access type of the oauth2.
type AuthOAuth2AccessType string

const (
	// AuthOAuth2AccessTypeOnline is the access type of the online.
	AuthOAuth2AccessTypeOnline AuthOAuth2AccessType = "online"
	// AuthOAuth2AccessTypeOffline is the access type of the offline.
	AuthOAuth2AccessTypeOffline AuthOAuth2AccessType = "offline"
)

// AuthOAuth2Config is the configuration for the oauth2.
type AuthOAuth2Config struct {
	GrantType    *string               `mapstructure:"grant_type"`
	ClientID     *string               `mapstructure:"client_id"`
	Scope        []string              `mapstructure:"scope"`
	ClientSecret string                `mapstructure:"client_secret"`
	AccessType   *string               `mapstructure:"access_type"`
	AuthURL      *string               `mapstructure:"auth_url"`
	TokenURL     *string               `mapstructure:"token_url"`
	Username     *string               `mapstructure:"username"`
	Password     *string               `mapstructure:"password"`
	Credential   *AuthCredentialConfig `mapstructure:"credential"`
}

// ValidAuthOAuth2Config represents the valid auth oauth2 configuration
type ValidAuthOAuth2Config struct {
	GrantType    AuthOAuth2GrantType
	ClientID     string
	Scope        []string
	ClientSecret string
	AccessType   AuthOAuth2AccessType
	AuthURL      string
	TokenURL     string
	Username     string
	Password     string
	Credential   ValidAuthCredentialConfig
}

// Validate validates the oauth2 configuration.
func (c AuthOAuth2Config) Validate() (ValidAuthOAuth2Config, error) {
	var valid ValidAuthOAuth2Config
	var err error
	if c.GrantType == nil {
		return ValidAuthOAuth2Config{}, ErrAuthOAuth2GrantTypeRequired
	}
	switch *c.GrantType {
	case string(AuthOAuth2GrantTypeAuthorizationCode):
		valid.GrantType = AuthOAuth2GrantTypeAuthorizationCode
		if c.AccessType == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2AccessTypeRequired
		}
		switch *c.AccessType {
		case string(AuthOAuth2AccessTypeOnline):
			valid.AccessType = AuthOAuth2AccessTypeOnline
		case string(AuthOAuth2AccessTypeOffline):
			valid.AccessType = AuthOAuth2AccessTypeOffline
		default:
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2AccessTypeInvalid
		}
		if c.AuthURL == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2AuthURLRequired
		}
		valid.AuthURL = *c.AuthURL
		if c.TokenURL == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2TokenURLRequired
		}
		valid.TokenURL = *c.TokenURL
	case string(AuthOAuth2GrantTypeClientCredentials):
		valid.GrantType = AuthOAuth2GrantTypeClientCredentials
		if c.TokenURL == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2TokenURLRequired
		}
		valid.TokenURL = *c.TokenURL
	case string(AuthOAuth2GrantTypePassword):
		valid.GrantType = AuthOAuth2GrantTypePassword
		if c.TokenURL == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2TokenURLRequired
		}
		valid.TokenURL = *c.TokenURL
		if c.Username == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2UsernameRequired
		}
		valid.Username = *c.Username
		if c.Password == nil {
			return ValidAuthOAuth2Config{}, ErrAuthOAuth2PasswordRequired
		}
		valid.Password = *c.Password
	default:
		return ValidAuthOAuth2Config{}, ErrAuthOAuth2GrantTypeInvalid
	}
	if c.ClientID == nil {
		return ValidAuthOAuth2Config{}, ErrAuthOAuth2ClientIDRequired
	}
	valid.ClientID = *c.ClientID
	valid.Scope = c.Scope
	valid.ClientSecret = c.ClientSecret

	if c.Credential == nil {
		return ValidAuthOAuth2Config{}, ErrAuthOAuth2CredentialRequired
	}
	if valid.Credential, err = c.Credential.Validate(); err != nil {
		return ValidAuthOAuth2Config{}, fmt.Errorf("credential: %w", err)
	}
	return valid, nil
}

// AuthAPIKeyConfig is the configuration for the api key.
type AuthAPIKeyConfig struct {
	HeaderName *string `mapstructure:"header_name"`
	Key        *string `mapstructure:"key"`
}

// ValidAuthAPIKeyConfig represents the valid auth api key configuration
type ValidAuthAPIKeyConfig struct {
	HeaderName string
	Key        string
}

// Validate validates the api key configuration.
func (c AuthAPIKeyConfig) Validate() (ValidAuthAPIKeyConfig, error) {
	var valid ValidAuthAPIKeyConfig
	if c.HeaderName == nil {
		return ValidAuthAPIKeyConfig{}, ErrAuthAPIKeyHeaderNameRequired
	}
	valid.HeaderName = *c.HeaderName
	if c.Key == nil {
		return ValidAuthAPIKeyConfig{}, ErrAuthAPIKeyKeyRequired
	}
	valid.Key = *c.Key
	return valid, nil
}

// AuthBasicConfig is the configuration for the basic auth.
type AuthBasicConfig struct {
	Username *string `mapstructure:"username"`
	Password *string `mapstructure:"password"`
}

// ValidAuthBasicConfig represents the valid auth basic configuration
type ValidAuthBasicConfig struct {
	Username string
	Password string
}

// Validate validates the basic auth configuration.
func (c AuthBasicConfig) Validate() (ValidAuthBasicConfig, error) {
	var valid ValidAuthBasicConfig
	if c.Username == nil {
		return ValidAuthBasicConfig{}, ErrAuthBasicUsernameRequired
	}
	valid.Username = *c.Username
	if c.Password == nil {
		return ValidAuthBasicConfig{}, ErrAuthBasicPasswordRequired
	}
	valid.Password = *c.Password
	return valid, nil
}

// AuthPrivateKeyConfig is the configuration for the private key.
type AuthPrivateKeyConfig struct {
	PrivateKey *string `mapstructure:"private_key"`
}

// ValidAuthPrivateKeyConfig represents the valid auth private key configuration
type ValidAuthPrivateKeyConfig struct {
	PrivateKey string
}

// Validate validates the private key configuration.
func (c AuthPrivateKeyConfig) Validate() (ValidAuthPrivateKeyConfig, error) {
	var valid ValidAuthPrivateKeyConfig
	if c.PrivateKey == nil {
		return ValidAuthPrivateKeyConfig{}, ErrAuthPrivateKeyPrivateKeyRequired
	}
	valid.PrivateKey = *c.PrivateKey
	return valid, nil
}

// AuthJWTConfig is the configuration for the jwt.
type AuthJWTConfig struct {
	Credential *AuthCredentialConfig `mapstructure:"credential"`
}

// ValidAuthJWTConfig represents the valid auth jwt configuration
type ValidAuthJWTConfig struct {
	Credential ValidAuthCredentialConfig
}

// Validate validates the jwt configuration.
func (c AuthJWTConfig) Validate() (ValidAuthJWTConfig, error) {
	var valid ValidAuthJWTConfig
	var err error
	if c.Credential == nil {
		return ValidAuthJWTConfig{}, ErrAuthJWTCredentialRequired
	}
	if valid.Credential, err = c.Credential.Validate(); err != nil {
		return ValidAuthJWTConfig{}, fmt.Errorf("credential: %w", err)
	}
	return valid, nil
}

// AuthConfig is the configuration for the auth.
type AuthConfig []AuthRespectiveConfig

// ValidAuthConfig represents the valid auth configuration
type ValidAuthConfig []ValidAuthRespectiveConfig

// Validate validates the auth configuration.
func (c AuthConfig) Validate() (ValidAuthConfig, error) {
	var valid ValidAuthConfig
	idSet := make(map[string]struct{})
	var hasDefault bool
	for i, ac := range c {
		var validRespective ValidAuthRespectiveConfig
		if ac.ID == nil {
			return ValidAuthConfig{}, fmt.Errorf("auth[%d].id: %w", i, ErrAuthIDRequired)
		}
		if _, ok := idSet[*ac.ID]; ok {
			return ValidAuthConfig{}, fmt.Errorf("auth[%d].id: %w", i, ErrAuthIDDuplicate)
		}
		idSet[*ac.ID] = struct{}{}
		validRespective.ID = *ac.ID
		if ac.Default {
			if hasDefault {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].default: %w", i, ErrAuthDefaultDuplicate)
			}
			hasDefault = true
		}
		validRespective.Default = ac.Default
		if ac.Type == nil {
			return ValidAuthConfig{}, fmt.Errorf("auth[%d].type: %w", i, ErrAuthTypeRequired)
		}
		switch *ac.Type {
		case string(AuthTypeOAuth2):
			if ac.OAuth2 == nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].oauth2: %w", i, ErrAuthOAuth2Required)
			}
			validOAuth2, err := ac.OAuth2.Validate()
			if err != nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].oauth2: %w", i, err)
			}
			validRespective.Type = AuthTypeOAuth2
			validRespective.OAuth2 = validOAuth2
		case string(AuthTypeAPIKey):
			if ac.APIKey == nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].apiKey: %w", i, ErrAuthAPIKeyRequired)
			}
			validAPIKey, err := ac.APIKey.Validate()
			if err != nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].apiKey: %w", i, err)
			}
			validRespective.Type = AuthTypeAPIKey
			validRespective.APIKey = validAPIKey
		case string(AuthTypeBasic):
			if ac.Basic == nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].basic: %w", i, ErrAuthBasicRequired)
			}
			validBasic, err := ac.Basic.Validate()
			if err != nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].basic: %w", i, err)
			}
			validRespective.Type = AuthTypeBasic
			validRespective.Basic = validBasic
		case string(AuthTypePrivateKey):
			if ac.PrivateKey == nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].privateKey: %w", i, ErrAuthPrivateKeyRequired)
			}
			validPrivateKey, err := ac.PrivateKey.Validate()
			if err != nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].privateKey: %w", i, err)
			}
			validRespective.Type = AuthTypePrivateKey
			validRespective.PrivateKey = validPrivateKey
		case string(AuthTypeJWT):
			if ac.JWT == nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].jwt: %w", i, ErrAuthJWTRequired)
			}
			validJWT, err := ac.JWT.Validate()
			if err != nil {
				return ValidAuthConfig{}, fmt.Errorf("auth[%d].jwt: %w", i, err)
			}
			validRespective.Type = AuthTypeJWT
			validRespective.JWT = validJWT
		default:
			return ValidAuthConfig{}, fmt.Errorf("auth[%d].type: %w", i, ErrAuthTypeInvalid)
		}
		valid = append(valid, validRespective)
	}
	return valid, nil
}
