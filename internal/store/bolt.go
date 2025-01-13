package store

import (
	"bytes"
	"fmt"
	"io"

	"github.com/boltdb/bolt"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/utils"
)

// BoltStore is a store that uses BoltDB
type BoltStore struct {
	db *bolt.DB
}

// SetupStore sets up the BoltStore
func (b *BoltStore) SetupStore(env string, conf config.ValidStoreConfig) error {
	for _, f := range conf.File {
		if f.Env == env {
			_, err := utils.CreateFileWithDir(f.Path)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			db, err := bolt.Open(f.Path, 0o600, &bolt.Options{
				// Timeout: 3 * time.Second,
			})
			if err != nil {
				return fmt.Errorf("failed to open bolt db: %w", err)
			}
			b.db = db
			if err := b.CreateBuckets(conf); err != nil {
				return fmt.Errorf("failed to create buckets: %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("no store config found for env %s", env)
}

// CreateBuckets creates buckets in the store
func (b *BoltStore) CreateBuckets(conf config.ValidStoreConfig) error {
	if err := b.db.Update(func(tx *bolt.Tx) error {
		for _, bucket := range conf.Buckets {
			if _, err := tx.CreateBucketIfNotExists([]byte(bucket)); err != nil {
				return fmt.Errorf("failed to create bucket: %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to create buckets: %w", err)
	}

	return nil
}

// PutObject puts an object in the store
func (b *BoltStore) PutObject(bucket, key string, data []byte) error {
	if err := b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		return b.Put([]byte(key), data)
	}); err != nil {
		return fmt.Errorf("failed to put object: %w", err)
	}

	return nil
}

// GetObject gets an object from the store
func (b *BoltStore) GetObject(bucket, key string) ([]byte, error) {
	var data []byte
	if err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		data = b.Get([]byte(key))
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	return data, nil
}

// PutObjectReader puts an object in the store from a reader
func (b *BoltStore) PutObjectReader(bucket, key string, reader io.Reader) error {
	if err := b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		data, err := io.ReadAll(reader)
		if err != nil {
			return fmt.Errorf("failed to read data: %w", err)
		}
		return b.Put([]byte(key), data)
	}); err != nil {
		return fmt.Errorf("failed to put object: %w", err)
	}

	return nil
}

// GetObjectReader gets an object from the store as a reader
func (b *BoltStore) GetObjectReader(bucket, key string) (io.Reader, error) {
	var data []byte
	if err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		data = b.Get([]byte(key))
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	return bytes.NewReader(data), nil
}

// DeleteObject deletes an object from the store
func (b *BoltStore) DeleteObject(bucket, key string) error {
	if err := b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		return b.Delete([]byte(key))
	}); err != nil {
		return fmt.Errorf("failed to delete object: %w", err)
	}

	return nil
}

// ListObjects lists all objects in a bucket
func (b *BoltStore) ListObjects(bucket string) ([]string, error) {
	var keys []string
	if err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s not found", bucket)
		}
		return b.ForEach(func(k, _ []byte) error {
			keys = append(keys, string(k))
			return nil
		})
	}); err != nil {
		return nil, fmt.Errorf("failed to list objects: %w", err)
	}

	return keys, nil
}

// ListBuckets lists all buckets in the store
func (b *BoltStore) ListBuckets() ([]string, error) {
	var buckets []string
	if err := b.db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			buckets = append(buckets, string(name))
			return nil
		})
	}); err != nil {
		return nil, fmt.Errorf("failed to list buckets: %w", err)
	}

	return buckets, nil
}

// Backup writes a backup of the store to a writer
func (b *BoltStore) Backup(writer io.Writer) (int, error) {
	var size int
	err := b.db.View(func(tx *bolt.Tx) error {
		if _, err := tx.WriteTo(writer); err != nil {
			return fmt.Errorf("failed to write backup: %w", err)
		}
		size = int(tx.Size())
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to backup store: %w", err)
	}

	return size, nil
}

// Clear clears the store
func (b *BoltStore) Clear() error {
	if err := b.db.Update(func(tx *bolt.Tx) error {
		if err := tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			if err := tx.DeleteBucket(name); err != nil {
				return fmt.Errorf("failed to delete bucket: %w", err)
			}

			return nil
		}); err != nil {
			return fmt.Errorf("failed to clear store: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("failed to clear store: %w", err)
	}

	return nil
}

// Close closes the store
func (b *BoltStore) Close() error {
	if err := b.db.Close(); err != nil {
		return fmt.Errorf("failed to close bolt db: %w", err)
	}

	return nil
}

var _ Store = &BoltStore{}
