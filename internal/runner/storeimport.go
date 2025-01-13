package runner

import (
	"context"
	"fmt"
	"sync"
)

// StoreImport represents the StoreImport runner
type StoreImport struct {
	Data []StoreImportData `yaml:"data"`
}

// ValidStoreImport represents the valid StoreImport runner
type ValidStoreImport struct {
	Data []ValidStoreImportData
}

// Validate validates the StoreImport
func (r StoreImport) Validate() (ValidStoreImport, error) {
	var validData []ValidStoreImportData
	for i, d := range r.Data {
		valid, err := d.Validate()
		if err != nil {
			return ValidStoreImport{}, fmt.Errorf("failed to validate data at index %d: %w", i, err)
		}
		validData = append(validData, valid)
	}
	return ValidStoreImport{
		Data: validData,
	}, nil
}

// StoreImportData represents the data for the StoreImport runner
type StoreImportData struct {
	BucketID   *string                 `yaml:"bucket_id"`
	Key        *string                 `yaml:"key"`
	StoreKey   *string                 `yaml:"store_key"`
	ThreadOnly bool                    `yaml:"thread_only"`
	Encrypt    CredentialEncryptConfig `yaml:"encrypt"`
}

// ValidStoreImportData represents the valid data for the StoreImport runner
type ValidStoreImportData struct {
	BucketID   string
	Key        string
	StoreKey   string
	ThreadOnly bool
	Encrypt    ValidCredentialEncryptConfig
}

// Validate validates the StoreImportData
func (d StoreImportData) Validate() (ValidStoreImportData, error) {
	if d.BucketID == nil {
		return ValidStoreImportData{}, fmt.Errorf("bucket_id is required")
	}
	if d.Key == nil {
		return ValidStoreImportData{}, fmt.Errorf("key is required")
	}
	if d.StoreKey == nil {
		return ValidStoreImportData{}, fmt.Errorf("store_key is required")
	}
	validEncrypt, err := d.Encrypt.Validate()
	if err != nil {
		return ValidStoreImportData{}, fmt.Errorf("failed to validate encrypt: %w", err)
	}
	return ValidStoreImportData{
		BucketID:   *d.BucketID,
		Key:        *d.Key,
		StoreKey:   *d.StoreKey,
		ThreadOnly: d.ThreadOnly,
		Encrypt:    validEncrypt,
	}, nil
}

// Run runs the StoreImport runner
func (r ValidStoreImport) Run(ctx context.Context, str Store, store *sync.Map) error {
	if err := str.Import(ctx, r.Data, func(_ context.Context, data ValidStoreImportData, val any, _ []byte) error {
		store.Store(data.Key, val)
		return nil
	}); err != nil {
		return fmt.Errorf("failed to import data: %w", err)
	}

	return nil
}
