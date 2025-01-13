package slave

import "fmt"

var (
	// ErrRequestNotFound is returned when the request is not found
	ErrRequestNotFound = fmt.Errorf("request not found")
	// ErrInvalidConnectionID represents an error when the connection ID is invalid
	ErrInvalidConnectionID = fmt.Errorf("invalid connection ID")
	// ErrInvalidEnvironment represents an error when the environment is invalid
	ErrInvalidEnvironment = fmt.Errorf("must connect to the same environment")
	// ErrInvalidCommandID represents an error when the command ID is invalid
	ErrInvalidCommandID = fmt.Errorf("invalid command ID")
	// ErrLoaderNotFound represents an error when the loader is not found
	ErrLoaderNotFound = fmt.Errorf("loader not found")
	// ErrFailedToSendLoaderResourceRequest represents an error when the loader resource request is failed to send
	ErrFailedToSendLoaderResourceRequest = fmt.Errorf("failed to send loader resource request")
	// ErrCommandNotFound represents an error when the command is not found
	ErrCommandNotFound = fmt.Errorf("command not found")
)
