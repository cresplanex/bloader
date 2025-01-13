package runner

import (
	"fmt"
)

const (
	// SlaveConnectRunnerEventConnecting represents the connecting event
	SlaveConnectRunnerEventConnecting Event = "slaveConnect:connecting"
	// SlaveConnectRunnerEventConnected represents the connected event
	SlaveConnectRunnerEventConnected Event = "slaveConnect:connected"
)

// SlaveConnect represents the SlaveConnect runner
type SlaveConnect struct {
	Slaves []SlaveConnectData `yaml:"slaves"`
}

// Validate validates the SlaveConnect
func (r SlaveConnect) Validate() (ValidSlaveConnect, error) {
	var validSlaves []ValidSlaveConnectData
	for i, d := range r.Slaves {
		valid, err := d.Validate()
		if err != nil {
			return ValidSlaveConnect{}, fmt.Errorf("failed to validate data at index %d: %w", i, err)
		}
		validSlaves = append(validSlaves, valid)
	}
	return ValidSlaveConnect{
		Slaves: validSlaves,
	}, nil
}

// SlaveConnectData represents the data for the SlaveConnect
type SlaveConnectData struct {
	ID          *string                 `yaml:"id"`
	URI         *string                 `yaml:"uri"`
	Certificate SlaveConnectCertificate `yaml:"certificate"`
	Encrypt     CredentialEncryptConfig `yaml:"encrypt"`
}

// Validate validates the SlaveConnectData
func (d SlaveConnectData) Validate() (ValidSlaveConnectData, error) {
	var valid ValidSlaveConnectData
	if d.ID == nil {
		return ValidSlaveConnectData{}, fmt.Errorf("id is required")
	}
	valid.ID = *d.ID
	if d.URI == nil {
		return ValidSlaveConnectData{}, fmt.Errorf("uri is required")
	}
	valid.URI = *d.URI
	validCertificate, err := d.Certificate.Validate()
	if err != nil {
		return ValidSlaveConnectData{}, fmt.Errorf("failed to validate certificate: %w", err)
	}
	valid.Certificate = validCertificate
	validEncrypt, err := d.Encrypt.Validate()
	if err != nil {
		return ValidSlaveConnectData{}, fmt.Errorf("failed to validate encrypt: %w", err)
	}
	valid.Encrypt = ValidCredentialEncryptConfig(validEncrypt)
	return valid, nil
}

// SlaveConnectCertificate represents the certificate for the Slave
type SlaveConnectCertificate struct {
	Enabled            bool    `yaml:"enabled"`
	CACert             *string `yaml:"ca_cert"`
	ServerNameOverride string  `yaml:"server_name_override"`
	InsecureSkipVerify bool    `yaml:"insecure_skip_verify"`
}

// Validate validates the SlaveConnectCertificate
func (c SlaveConnectCertificate) Validate() (ValidSlaveConnectCertificate, error) {
	if !c.Enabled {
		return ValidSlaveConnectCertificate{}, nil
	}
	if c.CACert == nil {
		return ValidSlaveConnectCertificate{}, fmt.Errorf("ca_cert is required")
	}
	return ValidSlaveConnectCertificate{
		Enabled:            c.Enabled,
		CACert:             *c.CACert,
		ServerNameOverride: c.ServerNameOverride,
		InsecureSkipVerify: c.InsecureSkipVerify,
	}, nil
}

// ValidSlaveConnect represents the valid ValidSlaveConnect runner
type ValidSlaveConnect struct {
	Slaves []ValidSlaveConnectData
}

// ValidSlaveConnectData represents the valid data for the ValidSlaveConnect
type ValidSlaveConnectData struct {
	ID          string
	URI         string
	Certificate ValidSlaveConnectCertificate
	Encrypt     ValidCredentialEncryptConfig
}

// ValidSlaveConnectCertificate represents the valid certificate for the Slave
type ValidSlaveConnectCertificate struct {
	Enabled            bool
	CACert             string
	ServerNameOverride string
	InsecureSkipVerify bool
}
