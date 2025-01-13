package config

import "fmt"

// StoreFileConfig represents the configuration for the file store.
type StoreFileConfig struct {
	Env  *string `mapstructure:"env"`
	Path *string `mapstructure:"path"`
}

// ValidStoreFileConfig represents the valid file store configuration
type ValidStoreFileConfig struct {
	Env  string
	Path string
}

// StoreConfig represents the configuration for the respective store.
type StoreConfig struct {
	File    []StoreFileConfig `mapstructure:"file"`
	Buckets []string          `mapstructure:"buckets"`
}

// ValidStoreConfig represents the valid store configuration
type ValidStoreConfig struct {
	File    []ValidStoreFileConfig
	Buckets []string
}

// Validate validates the store configuration.
func (c StoreConfig) Validate() (ValidStoreConfig, error) {
	var valid ValidStoreConfig
	for i, f := range c.File {
		var validFile ValidStoreFileConfig
		if f.Env == nil {
			return ValidStoreConfig{}, fmt.Errorf("store.file[%d].env: %w", i, ErrStoreFileEnvRequired)
		}
		validFile.Env = *f.Env
		if f.Path == nil {
			return ValidStoreConfig{}, fmt.Errorf("store.file[%d].path: %w", i, ErrStoreFilePathRequired)
		}
		validFile.Path = *f.Path
		valid.File = append(valid.File, validFile)
	}
	bucketIDSet := make(map[string]struct{})
	for i, b := range c.Buckets {
		if _, ok := bucketIDSet[b]; ok {
			return ValidStoreConfig{}, fmt.Errorf("store.buckets[%d]: %w", i, ErrStoreBucketIDDuplicate)
		}
		bucketIDSet[b] = struct{}{}
	}
	valid.Buckets = c.Buckets

	return valid, nil
}

// StoreSpecifyConfig represents the configuration for the store
type StoreSpecifyConfig struct {
	BucketID *string                 `mapstructure:"bucket_id"`
	Key      *string                 `mapstructure:"key"`
	Encrypt  CredentialEncryptConfig `mapstructure:"encrypt"`
}

// ValidStoreSpecifyConfig represents the valid store configuration
type ValidStoreSpecifyConfig struct {
	BucketID string
	Key      string
	Encrypt  ValidCredentialEncryptConfig
}

// CredentialEncryptConfig is the configuration for the credential encrypt.
type CredentialEncryptConfig struct {
	Enabled   bool    `mapstructure:"enabled"`
	EncryptID *string `mapstructure:"encrypt_id"`
}

// ValidCredentialEncryptConfig represents the valid auth credential encrypt configuration
type ValidCredentialEncryptConfig struct {
	Enabled   bool
	EncryptID string
}

// Validate validates the store configuration.
func (c StoreSpecifyConfig) Validate() (ValidStoreSpecifyConfig, error) {
	var valid ValidStoreSpecifyConfig
	if c.BucketID == nil {
		return ValidStoreSpecifyConfig{}, ErrStoreBucketIDRequired
	}
	valid.BucketID = *c.BucketID
	if c.Key == nil {
		return ValidStoreSpecifyConfig{}, ErrStoreKeyRequired
	}
	valid.Key = *c.Key
	if c.Encrypt.Enabled {
		if c.Encrypt.EncryptID == nil {
			return ValidStoreSpecifyConfig{}, ErrStoreEncryptIDRequired
		}
		valid.Encrypt.Enabled = c.Encrypt.Enabled
		valid.Encrypt.EncryptID = *c.Encrypt.EncryptID
	}
	return valid, nil
}
