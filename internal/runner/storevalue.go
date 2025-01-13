package runner

import (
	"context"
	"fmt"
)

// StoreValue represents the StoreValue runner
type StoreValue struct {
	Data []StoreValueData `yaml:"data"`
}

// ValidStoreValue represents the valid StoreValue runner
type ValidStoreValue struct {
	Data []ValidStoreValueData
}

// Validate validates the StoreValue
func (r StoreValue) Validate() (ValidStoreValue, error) {
	var validData []ValidStoreValueData
	for i, d := range r.Data {
		valid, err := d.Validate()
		if err != nil {
			return ValidStoreValue{}, fmt.Errorf("failed to validate data at index %d: %w", i, err)
		}
		validData = append(validData, valid)
	}
	return ValidStoreValue{
		Data: validData,
	}, nil
}

// StoreValueData represents the data for the StoreValue runner
type StoreValueData struct {
	BucketID *string                 `yaml:"bucket_id"`
	Key      *string                 `yaml:"key"`
	Value    *any                    `yaml:"value"`
	Encrypt  CredentialEncryptConfig `yaml:"encrypt"`
}

// ValidStoreValueData represents the valid data for the StoreValue runner
type ValidStoreValueData struct {
	BucketID string
	Key      string
	Value    any
	Encrypt  ValidCredentialEncryptConfig
}

// Validate validates the StoreValueData
func (d StoreValueData) Validate() (ValidStoreValueData, error) {
	if d.BucketID == nil {
		return ValidStoreValueData{}, fmt.Errorf("bucket_id is required")
	}
	if d.Key == nil {
		return ValidStoreValueData{}, fmt.Errorf("key is required")
	}
	if d.Value == nil {
		return ValidStoreValueData{}, fmt.Errorf("value is required")
	}
	validEncrypt, err := d.Encrypt.Validate()
	if err != nil {
		return ValidStoreValueData{}, fmt.Errorf("failed to validate encrypt: %w", err)
	}
	return ValidStoreValueData{
		BucketID: *d.BucketID,
		Key:      *d.Key,
		Value:    *d.Value,
		Encrypt:  validEncrypt,
	}, nil
}

// Run runs the StoreValue runner
func (r ValidStoreValue) Run(ctx context.Context, str Store) error {
	if err := str.Store(ctx, r.Data, nil); err != nil {
		return fmt.Errorf("failed to store data: %w", err)
	}
	return nil
}
