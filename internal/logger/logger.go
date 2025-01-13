// Package logger provides logging functionality for the application
package logger

import (
	"context"

	"github.com/cresplanex/bloader/internal/config"
)

// Logger is an interface for logging
type Logger interface {
	SetupLogger(env string, conf config.ValidLoggingConfig) error
	With(args ...KeyVal) Logger
	Debug(ctx context.Context, msg string, args ...KeyVal)
	Info(ctx context.Context, msg string, args ...KeyVal)
	Warn(ctx context.Context, msg string, args ...KeyVal)
	Error(ctx context.Context, msg string, args ...KeyVal)
	Close() error
}

// NewLoggerFromConfig creates a new Logger from the config
func NewLoggerFromConfig(_ string, _ config.ValidLoggingConfig) (Logger, error) {
	return &SlogLogger{}, nil
}

// KeyVal is a key value pair
type KeyVal struct {
	Key   string
	Value any
}

// Value creates a new KeyVal
func Value(key string, value any) KeyVal {
	return KeyVal{Key: key, Value: value}
}

// Group creates a new KeyVal with a group of KeyVals
func Group(key string, kvs ...KeyVal) KeyVal {
	return KeyVal{Key: key, Value: kvs}
}
