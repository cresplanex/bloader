package slcontainer

import (
	"encoding/json"
	"fmt"
	"sync"
)

// StoreDataKey is the struct to store the store data key
type StoreDataKey struct {
	BucketID string
	StoreKey string
}

// Store is the interface to store the data
type Store struct {
	mu   *sync.RWMutex
	data map[StoreDataKey]any
}

// NewStore creates a new store
func NewStore() *Store {
	return &Store{
		mu:   &sync.RWMutex{},
		data: make(map[StoreDataKey]any),
	}
}

// AddData adds the data to the store
func (s *Store) AddData(bucketID, storeKey string, data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	s.data[StoreDataKey{BucketID: bucketID, StoreKey: storeKey}] = v

	return nil
}

// RemoveData removes the data from the store
func (s *Store) RemoveData(bucketID, storeKey string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, StoreDataKey{BucketID: bucketID, StoreKey: storeKey})
}

// GetData gets the data from the store
func (s *Store) GetData(bucketID, storeKey string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, ok := s.data[StoreDataKey{BucketID: bucketID, StoreKey: storeKey}]

	return data, ok
}

// ExistData checks if the data exists in the store
func (s *Store) ExistData(bucketID, storeKey string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[StoreDataKey{BucketID: bucketID, StoreKey: storeKey}]
	return ok
}

// StoreRespectiveRequest is the struct to store the store respective request
type StoreRespectiveRequest struct {
	BucketID   string
	StoreKey   string
	Encryption Encryption
}

// StoreResourceRequest is the struct to store the store resource request
type StoreResourceRequest struct {
	Requests []StoreRespectiveRequest
}

// StoreData is the struct to store the store data
type StoreData struct {
	BucketID   string
	StoreKey   string
	Data       []byte
	Encryption Encryption
}

// StoreDataRequest is the struct to store the store data request
type StoreDataRequest struct {
	StoreData []StoreData
}
