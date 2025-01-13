package slcontainer

import (
	"fmt"
	"strings"
	"sync"
)

// Loader represents the loader container for the slave node
type Loader struct {
	mu               *sync.RWMutex
	LoaderBuilderMap map[string]*strings.Builder
	LoaderMap        map[string]string
}

// NewLoader creates a new loader container for the slave node
func NewLoader() *Loader {
	return &Loader{
		mu:               &sync.RWMutex{},
		LoaderBuilderMap: make(map[string]*strings.Builder),
		LoaderMap:        make(map[string]string),
	}
}

// WriteString writes a string to the loader container
func (l *Loader) WriteString(loaderID, data string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.LoaderBuilderMap[loaderID]; !ok {
		l.LoaderBuilderMap[loaderID] = &strings.Builder{}
	}
	if _, err := l.LoaderBuilderMap[loaderID].WriteString(data); err != nil {
		return fmt.Errorf("failed to write string: %w", err)
	}

	return nil
}

// Build builds the loader container
func (l *Loader) Build(loaderID string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.LoaderBuilderMap[loaderID]; ok {
		l.LoaderMap[loaderID] = l.LoaderBuilderMap[loaderID].String()
		delete(l.LoaderBuilderMap, loaderID)
	}
}

// GetLoader returns the loader from the container
func (l *Loader) GetLoader(loaderID string) (string, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if loader, ok := l.LoaderMap[loaderID]; ok {
		return loader, true
	}
	return "", false
}

// LoaderResourceRequest is a struct that represents a request to the loader resource.
type LoaderResourceRequest struct {
	LoaderID string
}
