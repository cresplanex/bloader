// Package config provides configuration for the application.
package config

// Type represents the configuration type
type Type string

const (
	// ConfigTypeMaster represents the master configuration type
	ConfigTypeMaster Type = "master"
	// ConfigTypeSlave represents the slave configuration type
	ConfigTypeSlave Type = "slave"
)

// ForOverride represents the configuration for the override service
type ForOverride struct {
	Env      *string         `mapstructure:"env"`
	Override *OverrideConfig `mapstructure:"override"`
}

// ValidConfigForOverride represents the configuration for the override service
type ValidConfigForOverride struct {
	Env      string              `mapstructure:"env"`
	Override ValidOverrideConfig `mapstructure:"override"`
}

// Validate validates the configuration for the override service
func (c ForOverride) Validate() (ValidConfigForOverride, error) {
	var valid ValidConfigForOverride
	if c.Env == nil {
		return ValidConfigForOverride{}, ErrEnvRequired
	}
	valid.Env = *c.Env
	if c.Override == nil {
		return ValidConfigForOverride{}, ErrOverrideRequired
	}
	validOverride, err := c.Override.Validate()
	if err != nil {
		return ValidConfigForOverride{}, err
	}
	valid.Override = validOverride

	return valid, nil
}

// Config represents the application configuration
type Config struct {
	Type         *string             `mapstructure:"type"`
	Env          *string             `mapstructure:"env"`
	Loader       *LoaderConfig       `mapstructure:"loader"`
	Targets      *TargetConfig       `mapstructure:"targets"`
	Outputs      *OutputConfig       `mapstructure:"outputs"`
	Store        *StoreConfig        `mapstructure:"store"`
	Encrypts     *EncryptConfig      `mapstructure:"encrypts"`
	Auth         *AuthConfig         `mapstructure:"auth"`
	Server       *ServerConfig       `mapstructure:"server"`
	Logging      *LoggingConfig      `mapstructure:"logging"`
	Clock        *ClockConfig        `mapstructure:"clock"`
	Language     *LanguageConfig     `mapstructure:"language"`
	Override     *OverrideConfig     `mapstructure:"override"`
	SlaveSetting *SlaveSettingConfig `mapstructure:"slave_setting"`
}

// ValidConfig represents the application configuration
type ValidConfig struct {
	Type         Type
	Env          string
	Loader       ValidLoaderConfig
	Targets      ValidTargetConfig
	Outputs      ValidOutputConfig
	Store        ValidStoreConfig
	Encrypts     ValidEncryptConfig
	Auth         ValidAuthConfig
	Server       ValidServerConfig
	Logging      ValidLoggingConfig
	Clock        ValidClockConfig
	Language     ValidLanguageConfig
	Override     ValidOverrideConfig
	SlaveSetting ValidSlaveSettingConfig
}

// Validate validates the configuration
func (c Config) Validate() (ValidConfig, error) {
	var valid ValidConfig
	if c.Type == nil {
		return ValidConfig{}, ErrTypeRequired
	}
	switch *c.Type {
	case string(ConfigTypeMaster):
		valid.Type = ConfigTypeMaster
		err := c.validateMaster(&valid)
		if err != nil {
			return ValidConfig{}, err
		}
	case string(ConfigTypeSlave):
		valid.Type = ConfigTypeSlave
		err := c.validateSlave(&valid)
		if err != nil {
			return ValidConfig{}, err
		}
	default:
		return ValidConfig{}, ErrTypeInvalid
	}

	return valid, nil
}

func (c Config) validateMaster(valid *ValidConfig) error {
	if c.Env == nil {
		return ErrEnvRequired
	}
	valid.Env = *c.Env

	if c.Loader == nil {
		return ErrLoaderRequired
	}
	validLoader, err := c.Loader.Validate()
	if err != nil {
		return err
	}
	valid.Loader = validLoader

	if c.Targets == nil {
		return ErrTargetsRequired
	}
	validTargets, err := c.Targets.Validate()
	if err != nil {
		return err
	}
	valid.Targets = validTargets

	if c.Outputs == nil {
		return ErrOutputsRequired
	}
	validOutputs, err := c.Outputs.Validate()
	if err != nil {
		return err
	}
	valid.Outputs = validOutputs

	if c.Store == nil {
		return ErrStoreRequired
	}
	validStore, err := c.Store.Validate()
	if err != nil {
		return err
	}
	valid.Store = validStore

	if c.Encrypts == nil {
		return ErrEncryptsRequired
	}
	validEncrypts, err := c.Encrypts.Validate()
	if err != nil {
		return err
	}
	valid.Encrypts = validEncrypts

	if c.Auth == nil {
		return ErrAuthRequired
	}
	validAuth, err := c.Auth.Validate()
	if err != nil {
		return err
	}
	valid.Auth = validAuth

	if c.Server == nil {
		return ErrServerRequired
	}
	validServer, err := c.Server.Validate()
	if err != nil {
		return err
	}
	valid.Server = validServer

	if c.Logging == nil {
		return ErrLoggingRequired
	}
	validLogging, err := c.Logging.Validate()
	if err != nil {
		return err
	}
	valid.Logging = validLogging

	if c.Clock == nil {
		return ErrClockRequired
	}
	validClock, err := c.Clock.Validate()
	if err != nil {
		return err
	}
	valid.Clock = validClock

	if c.Language == nil {
		return ErrLanguageRequired
	}
	validLanguage, err := c.Language.Validate()
	if err != nil {
		return err
	}
	valid.Language = validLanguage

	if c.Override == nil {
		return ErrOverrideRequired
	}
	validOverride, err := c.Override.Validate()
	if err != nil {
		return err
	}
	valid.Override = validOverride

	return nil
}

func (c Config) validateSlave(valid *ValidConfig) error {
	if c.Env == nil {
		return ErrEnvRequired
	}
	valid.Env = *c.Env

	if c.Encrypts == nil {
		return ErrEncryptsRequired
	}
	validEncrypts, err := c.Encrypts.ValidateOnSlave()
	if err != nil {
		return err
	}
	valid.Encrypts = validEncrypts

	if c.Logging == nil {
		return ErrLoggingRequired
	}
	validLogging, err := c.Logging.Validate()
	if err != nil {
		return err
	}
	valid.Logging = validLogging

	if c.Clock == nil {
		return ErrClockRequired
	}
	validClock, err := c.Clock.Validate()
	if err != nil {
		return err
	}
	valid.Clock = validClock

	if c.Language == nil {
		return ErrLanguageRequired
	}
	validLanguage, err := c.Language.Validate()
	if err != nil {
		return err
	}
	valid.Language = validLanguage

	if c.Override == nil {
		return ErrOverrideRequired
	}
	validOverride, err := c.Override.Validate()
	if err != nil {
		return err
	}
	valid.Override = validOverride

	if c.SlaveSetting == nil {
		return ErrSlaveSettingRequired
	}
	validSlaveSetting, err := c.SlaveSetting.Validate()
	if err != nil {
		return err
	}
	valid.SlaveSetting = validSlaveSetting

	return nil
}
