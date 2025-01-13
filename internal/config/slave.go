package config

// SlaveSettingConfig represents the configuration for the slave setting
type SlaveSettingConfig struct {
	Port        *int                    `mapstructure:"port"`
	Certificate SlaveCertificateConfig  `mapstructure:"certificate"`
	Encrypt     CredentialEncryptConfig `mapstructure:"encrypt"`
}

// ValidSlaveSettingConfig represents the valid slave setting configuration
type ValidSlaveSettingConfig struct {
	Port        int
	Certificate ValidSlaveCertificateConfig
	Encrypt     ValidCredentialEncryptConfig
}

// Validate validates the slave setting configuration.
func (c SlaveSettingConfig) Validate() (ValidSlaveSettingConfig, error) {
	var valid ValidSlaveSettingConfig
	var err error
	if c.Port == nil {
		return ValidSlaveSettingConfig{}, ErrSlaveSettingPortRequired
	}
	valid.Port = *c.Port
	valid.Certificate, err = c.Certificate.Validate()
	if err != nil {
		return ValidSlaveSettingConfig{}, err
	}
	if c.Encrypt.Enabled {
		if c.Encrypt.EncryptID == nil {
			return ValidSlaveSettingConfig{}, ErrSlaveSettingEncryptIDRequired
		}
		valid.Encrypt.Enabled = c.Encrypt.Enabled
		valid.Encrypt.EncryptID = *c.Encrypt.EncryptID
	}
	return valid, nil
}

// SlaveCertificateConfig represents the configuration for the slave certificate
type SlaveCertificateConfig struct {
	Enabled   bool    `mapstructure:"enabled"`
	SlaveCert *string `mapstructure:"slave_cert"`
	SlaveKey  *string `mapstructure:"slave_key"`
}

// ValidSlaveCertificateConfig represents the valid slave certificate configuration
type ValidSlaveCertificateConfig struct {
	Enabled   bool
	SlaveCert string
	SlaveKey  string
}

// Validate validates the slave certificate configuration.
func (c SlaveCertificateConfig) Validate() (ValidSlaveCertificateConfig, error) {
	var valid ValidSlaveCertificateConfig
	if c.Enabled {
		if c.SlaveCert == nil {
			return ValidSlaveCertificateConfig{}, ErrSlaveCertificateSlaveCertPathRequired
		}
		valid.SlaveCert = *c.SlaveCert
		if c.SlaveKey == nil {
			return ValidSlaveCertificateConfig{}, ErrSlaveCertificateSlaveKeyPathRequired
		}
		valid.SlaveKey = *c.SlaveKey
	}
	valid.Enabled = c.Enabled
	return valid, nil
}
