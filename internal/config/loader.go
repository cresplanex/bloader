package config

// LoaderConfig represents the configuration for the loader service
type LoaderConfig struct {
	BasePath *string `mapstructure:"base_path"`
}

// ValidLoaderConfig represents the configuration for the loader service
type ValidLoaderConfig struct {
	BasePath string
}

// Validate validates the loader configuration
func (c LoaderConfig) Validate() (ValidLoaderConfig, error) {
	var valid ValidLoaderConfig
	if c.BasePath == nil {
		return ValidLoaderConfig{}, ErrLoaderBasePathRequired
	}
	valid.BasePath = *c.BasePath

	return valid, nil
}
