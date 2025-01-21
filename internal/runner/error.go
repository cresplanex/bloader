package runner

import "errors"

type syncError struct {
	Err error
}

var errByContext = errors.New("context is done")
