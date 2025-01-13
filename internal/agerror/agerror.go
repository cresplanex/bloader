// Package agerror provides aggregated error handling.
package agerror

import (
	"errors"
	"fmt"
	"strings"
)

// AggregateError is an error type that aggregates multiple errors.
type AggregateError struct {
	Errors []error
	prefix string
}

// NewAggregateErrorWithErr creates a new AggregateError.
func NewAggregateErrorWithErr(initErr error, prefix string) AggregateError {
	return AggregateError{
		[]error{
			initErr,
		},
		prefix,
	}
}

// NewAggregateError creates a new AggregateError.
func NewAggregateError(prefix string) AggregateError {
	return AggregateError{
		[]error{},
		prefix,
	}
}

// Error adds an error to the AggregateError.
func (e AggregateError) Error() string {
	var sb strings.Builder

	sb.WriteString(e.prefix + ":")

	for _, err := range e.Errors {
		sb.WriteString(fmt.Sprintf(" %v;", err.Error()))
	}

	return strings.TrimRight(sb.String(), ";")
}

// Is conforms with errors.Is.
func (e AggregateError) Is(err error) bool {
	for _, ie := range e.Errors {
		if errors.Is(ie, err) {
			return true
		}
	}
	return false
}

// Add adds an error to the AggregateError.
func (e *AggregateError) Add(err error) {
	e.Errors = append(e.Errors, err)
}
