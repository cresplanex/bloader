package slcontainer

import "fmt"

var (
	// ErrInvalidLoaderID is returned when the loader id is invalid
	ErrInvalidLoaderID = fmt.Errorf("invalid loader id")
	// ErrInvalidAuthID is returned when the auth id is invalid
	ErrInvalidAuthID = fmt.Errorf("invalid auth id")
	// ErrRequestNotFound is returned when the request is not found
	ErrRequestNotFound = fmt.Errorf("request not found")
	// ErrInvalidAuthType is returned when the auth type is invalid
	ErrInvalidAuthType = fmt.Errorf("invalid auth type")
)
