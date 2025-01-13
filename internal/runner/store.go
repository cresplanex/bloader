package runner

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cresplanex/bloader/internal/encrypt"
	"github.com/cresplanex/bloader/internal/store"
)

// StoreCallback represents the store callback
type StoreCallback func(ctx context.Context, data ValidStoreValueData, valBytes []byte) error

// StoreWithExtractorCallback represents the store with extractor callback
type StoreWithExtractorCallback func(ctx context.Context, data ValidExecRequestStoreData, valBytes []byte) error

// ImportCallback represents the import callback
// there are the case valBytes is nil
type ImportCallback func(ctx context.Context, data ValidStoreImportData, val any, valBytes []byte) error

// Store represents the store
type Store interface {
	// Store stores the data
	Store(ctx context.Context, data []ValidStoreValueData, cb StoreCallback) error
	// StoreWithExtractor stores the data with extractor
	StoreWithExtractor(ctx context.Context, res any, data []ValidExecRequestStoreData, cb StoreWithExtractorCallback) error
	// Import loads the data
	Import(ctx context.Context, data []ValidStoreImportData, cb ImportCallback) error
}

// LocalStore represents the local store
type LocalStore struct {
	encCtr encrypt.Container
	str    store.Store
}

// NewLocalStore creates a new local store
func NewLocalStore(encCtr encrypt.Container, str store.Store) *LocalStore {
	return &LocalStore{
		encCtr: encCtr,
		str:    str,
	}
}

// Store stores the data
func (l LocalStore) Store(ctx context.Context, data []ValidStoreValueData, cb StoreCallback) error {
	for _, d := range data {
		valBytes, err := json.Marshal(d.Value)
		if err != nil {
			return fmt.Errorf("failed to marshal store data: %w", err)
		}
		if d.Encrypt.Enabled {
			encrypter, ok := l.encCtr[d.Encrypt.EncryptID]
			if !ok {
				return fmt.Errorf("encrypter not found: %s", d.Encrypt.EncryptID)
			}
			encryptedVal, err := encrypter.Encrypt(valBytes)
			if err != nil {
				return fmt.Errorf("failed to encrypt value: %w", err)
			}
			valBytes = []byte(encryptedVal)
		}
		if cb != nil {
			if err := cb(ctx, d, valBytes); err != nil {
				return fmt.Errorf("failed to store data: %w", err)
			}
		}
		if err := l.str.PutObject(d.BucketID, d.Key, valBytes); err != nil {
			return fmt.Errorf("failed to put object: %w", err)
		}
	}
	return nil
}

// StoreWithExtractor stores the data with extractor
func (l LocalStore) StoreWithExtractor(
	ctx context.Context,
	res any,
	data []ValidExecRequestStoreData,
	cb StoreWithExtractorCallback,
) error {
	for _, d := range data {
		result, err := d.Extractor.Extract(res)
		if err != nil {
			return fmt.Errorf("failed to extract store data: %w", err)
		}
		valBytes, err := json.Marshal(result)
		if err != nil {
			return fmt.Errorf("failed to marshal store data: %w", err)
		}
		if d.Encrypt.Enabled {
			encrypter, ok := l.encCtr[d.Encrypt.EncryptID]
			if !ok {
				return fmt.Errorf("encrypter not found: %s", d.Encrypt.EncryptID)
			}
			encryptedVal, err := encrypter.Encrypt(valBytes)
			if err != nil {
				return fmt.Errorf("failed to encrypt value: %w", err)
			}
			valBytes = []byte(encryptedVal)
		}
		if cb != nil {
			if err := cb(ctx, d, valBytes); err != nil {
				return fmt.Errorf("failed to store data: %w", err)
			}
		}
		if err := l.str.PutObject(d.BucketID, d.StoreKey, valBytes); err != nil {
			return fmt.Errorf("failed to put object: %w", err)
		}
	}
	return nil
}

// Import loads the data
func (l LocalStore) Import(ctx context.Context, data []ValidStoreImportData, cb ImportCallback) error {
	for _, d := range data {
		valBytes, err := l.str.GetObject(d.BucketID, d.StoreKey)
		if err != nil {
			return fmt.Errorf("failed to get object: %w", err)
		}
		if len(valBytes) == 0 {
			return fmt.Errorf("object not found: %s in bucket: %s", d.StoreKey, d.BucketID)
		}
		if d.Encrypt.Enabled {
			encrypter, ok := l.encCtr[d.Encrypt.EncryptID]
			if !ok {
				return fmt.Errorf("encrypter not found: %s", d.Encrypt.EncryptID)
			}
			decryptedVal, err := encrypter.Decrypt(string(valBytes))
			if err != nil {
				return fmt.Errorf("failed to decrypt value: %w", err)
			}
			valBytes = []byte(decryptedVal)
		}
		var val any
		if err := json.Unmarshal(valBytes, &val); err != nil {
			return fmt.Errorf("failed to unmarshal value: %w", err)
		}
		if cb != nil {
			if err := cb(ctx, d, val, valBytes); err != nil {
				return fmt.Errorf("failed to import data: %w", err)
			}
		}
	}

	return nil
}

var _ Store = (*LocalStore)(nil)
