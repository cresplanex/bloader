package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFileWithDir creates a file with the directory
//
//nolint:gosec
func CreateFileWithDir(filePath string) (*os.File, error) {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	}
	var f *os.File
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err = os.Create(filepath.Clean(filePath))
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
	} else {
		f, err = os.Open(filepath.Clean(filePath))
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
	}
	return f, nil
}
