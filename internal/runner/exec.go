package runner

import (
	"fmt"

	"github.com/cresplanex/bloader/internal/runner/matcher"
)

// ExecRequestData represents the data configuration for the OneExec runner
type ExecRequestData struct {
	Key       *string                `yaml:"key"`
	Extractor *matcher.DataExtractor `yaml:"extractor"`
}

// ValidExecRequestData represents the valid data configuration for the OneExec runner
type ValidExecRequestData struct {
	Key       string
	Extractor matcher.ValidDataExtractor
}

// Validate validates the OneExecRequestData
func (d ExecRequestData) Validate() (ValidExecRequestData, error) {
	if d.Key == nil {
		return ValidExecRequestData{}, fmt.Errorf("key is required")
	}
	if d.Extractor == nil {
		return ValidExecRequestData{}, fmt.Errorf("extractor is required")
	}
	validExtractor, err := d.Extractor.Validate()
	if err != nil {
		return ValidExecRequestData{}, fmt.Errorf("failed to validate extractor: %w", err)
	}
	return ValidExecRequestData{
		Key:       *d.Key,
		Extractor: validExtractor,
	}, nil
}

// ValidExecRequestDataSlice represents a slice of ValidExecRequestData
type ValidExecRequestDataSlice []ValidExecRequestData

// ExtractHeader extracts the header from the data
func (s ValidExecRequestDataSlice) ExtractHeader() []string {
	var header []string
	for _, d := range s {
		header = append(header, d.Key)
	}
	return header
}

// ExecRequestStoreData represents the store data configuration for the OneExec runner
type ExecRequestStoreData struct {
	BucketID  *string                 `yaml:"bucket_id"`
	StoreKey  *string                 `yaml:"store_key"`
	Encrypt   CredentialEncryptConfig `yaml:"encrypt"`
	Extractor *matcher.DataExtractor  `yaml:"extractor"`
}

// ValidExecRequestStoreData represents the valid store data configuration for the OneExec runner
type ValidExecRequestStoreData struct {
	BucketID  string
	StoreKey  string
	Encrypt   ValidCredentialEncryptConfig
	Extractor matcher.ValidDataExtractor
}

// Validate validates the OneExecRequestStoreData
func (d ExecRequestStoreData) Validate() (ValidExecRequestStoreData, error) {
	if d.BucketID == nil {
		return ValidExecRequestStoreData{}, fmt.Errorf("bucket_id is required")
	}
	if d.StoreKey == nil {
		return ValidExecRequestStoreData{}, fmt.Errorf("store_key is required")
	}
	validEncrypt, err := d.Encrypt.Validate()
	if err != nil {
		return ValidExecRequestStoreData{}, fmt.Errorf("failed to validate encrypt: %w", err)
	}
	validExtractor, err := d.Extractor.Validate()
	if err != nil {
		return ValidExecRequestStoreData{}, fmt.Errorf("failed to validate extractor: %w", err)
	}
	return ValidExecRequestStoreData{
		BucketID:  *d.BucketID,
		StoreKey:  *d.StoreKey,
		Encrypt:   validEncrypt,
		Extractor: validExtractor,
	}, nil
}
