// Package utils provides utility functions for the application
package utils

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

// GenerateUniqueID generates a unique ID
func GenerateUniqueID() string {
	return uuid.New().String()
}

// GenerateRandomString generates a random string of the given length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GenerateRandomStringWithCharset generates a random string of the given length with the given charset.
func GenerateRandomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateRandomBytes generates a random byte slice of the given size.
func GenerateRandomBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := io.ReadFull(cryptorand.Reader, bytes); err != nil {
		return nil, fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return bytes, nil
}

// Contains checks if a slice contains a specific element
func Contains[T comparable](slice []T, elem T) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

// AnyContains checks if any of the slices contains a specific element
func AnyContains[T comparable](slices1, slices2 []T) bool {
	for _, e1 := range slices1 {
		for _, e2 := range slices2 {
			if e1 == e2 {
				return true
			}
		}
	}
	return false
}

// AllContains checks if all of the slices contain a specific element
func AllContains[T comparable](slices1, slices2 []T) bool {
	for _, e2 := range slices2 {
		found := false
		for _, e1 := range slices1 {
			if e1 == e2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// RemoveElement removes an element from a slice
func RemoveElement[T comparable](slice []T, elem T) []T {
	var result []T
	for _, e := range slice {
		if e != elem {
			result = append(result, e)
		}
	}
	return result
}
