package runner

import (
	"context"
	"fmt"
	"sync"
)

// MemoryValue represents the MemoryValue runner
type MemoryValue struct {
	Data []MemoryValueData `yaml:"data"`
}

// ValidMemoryValue represents the valid MemoryValue runner
type ValidMemoryValue struct {
	Data []ValidMemoryValueData
}

// Validate validates the MemoryValue
func (r MemoryValue) Validate() (ValidMemoryValue, error) {
	var validData []ValidMemoryValueData
	for i, d := range r.Data {
		valid, err := d.Validate()
		if err != nil {
			return ValidMemoryValue{}, fmt.Errorf("failed to validate data at index %d: %w", i, err)
		}
		validData = append(validData, valid)
	}
	return ValidMemoryValue{
		Data: validData,
	}, nil
}

// MemoryValueData represents the data for the MemoryValue runner
type MemoryValueData struct {
	Key   *string `yaml:"key"`
	Value *any    `yaml:"value"`
}

// ValidMemoryValueData represents the valid data for the MemoryValue runner
type ValidMemoryValueData struct {
	Key   string
	Value any
}

// Validate validates the MemoryValueData
func (d MemoryValueData) Validate() (ValidMemoryValueData, error) {
	if d.Key == nil {
		return ValidMemoryValueData{}, fmt.Errorf("key is required")
	}
	if d.Value == nil {
		return ValidMemoryValueData{}, fmt.Errorf("value is required")
	}
	return ValidMemoryValueData{
		Key:   *d.Key,
		Value: *d.Value,
	}, nil
}

// Run runs the MemoryValue runner
func (r ValidMemoryValue) Run(_ context.Context, store *sync.Map) error {
	for _, d := range r.Data {
		store.Store(d.Key, d.Value)
	}
	return nil
}
