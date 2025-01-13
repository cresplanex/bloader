package config

// ServerConfig represents the server configuration
type ServerConfig struct {
	Port         *int `mapstructure:"port"`
	RedirectPort *int `mapstructure:"redirect_port"`
}

// ValidServerConfig represents the valid server configuration
type ValidServerConfig struct {
	Port         int
	RedirectPort struct {
		Enabled bool
		Port    int
	}
}

// Validate validates the server configuration
func (s ServerConfig) Validate() (ValidServerConfig, error) {
	var valid ValidServerConfig
	if s.Port == nil {
		return ValidServerConfig{}, ErrServerPortRequired
	}
	valid.Port = *s.Port
	if s.RedirectPort != nil {
		valid.RedirectPort.Enabled = true
		valid.RedirectPort.Port = *s.RedirectPort
	}
	return valid, nil
}
