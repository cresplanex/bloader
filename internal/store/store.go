// Package store provides the store logic for the application
package store

import (
	"io"

	"github.com/cresplanex/bloader/internal/config"
)

// Store is the interface for a store
type Store interface {
	SetupStore(env string, conf config.ValidStoreConfig) error
	CreateBuckets(conf config.ValidStoreConfig) error
	PutObject(bucket, key string, data []byte) error
	GetObject(bucket, key string) ([]byte, error)
	PutObjectReader(bucket, key string, reader io.Reader) error
	GetObjectReader(bucket, key string) (io.Reader, error)
	DeleteObject(bucket, key string) error
	ListObjects(bucket string) ([]string, error)
	ListBuckets() ([]string, error)
	Backup(writer io.Writer) (int, error)
	Clear() error
	Close() error
}

// NewStoreFromConfig creates a new store from the configuration
func NewStoreFromConfig(_ config.ValidStoreConfig) (Store, error) {
	return &BoltStore{}, nil
}
