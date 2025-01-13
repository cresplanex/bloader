package logger

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"

	slogmulti "github.com/samber/slog-multi"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/utils"
)

// SlogLogger is a logger that uses slog
type SlogLogger struct {
	Logger    *slog.Logger
	openFIles []*os.File
}

// NewSlogLogger creates a new SlogLogger
func NewSlogLogger() *SlogLogger {
	return &SlogLogger{
		Logger: slog.Default(),
	}
}

// SetupLogger sets up the logger with the given configuration
func (l *SlogLogger) SetupLogger(env string, cfg config.ValidLoggingConfig) error {
	var handlers []slog.Handler

	for _, output := range cfg.Output {
		if !output.EnabledEnv.All && !utils.Contains(output.EnabledEnv.Values, env) {
			continue
		}

		var level slog.Level
		switch output.Level {
		case config.LoggingOutputLevelDebug:
			level = slog.LevelDebug
		case config.LoggingOutputLevelInfo:
			level = slog.LevelInfo
		case config.LoggingOutputLevelWarn:
			level = slog.LevelWarn
		case config.LoggingOutputLevelError:
			level = slog.LevelError
		default:
			level = slog.LevelInfo
		}

		var handler slog.Handler
		switch output.Type {
		case config.LoggingOutputTypeStdout:
			switch output.Format {
			case config.LoggingOutputFormatText:
				handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
			case config.LoggingOutputFormatJSON:
				handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
			}
		case config.LoggingOutputTypeFile:
			//nolint:gosec
			if err := os.MkdirAll(filepath.Dir(output.Filename), 0o755); err != nil {
				return fmt.Errorf("failed to create log directories: %w", err)
			}
			file, err := os.OpenFile(output.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
			if err != nil {
				return fmt.Errorf("failed to open log file: %w", err)
			}
			l.openFIles = append(l.openFIles, file)
			switch output.Format {
			case config.LoggingOutputFormatText:
				handler = slog.NewTextHandler(file, &slog.HandlerOptions{Level: level})
			case config.LoggingOutputFormatJSON:
				handler = slog.NewJSONHandler(file, &slog.HandlerOptions{Level: level})
			}
		case config.LoggingOutputTypeTCP:
			conn, err := net.Dial("tcp", output.Address)
			if err != nil {
				return fmt.Errorf("failed to connect to tcp server: %w", err)
			}
			switch output.Format {
			case config.LoggingOutputFormatText:
				handler = slog.NewTextHandler(conn, &slog.HandlerOptions{Level: level})
			case config.LoggingOutputFormatJSON:
				handler = slog.NewJSONHandler(conn, &slog.HandlerOptions{Level: level})
			}
		}

		if handler != nil {
			handlers = append(handlers, handler)
		}
	}
	logger := slog.New(
		slogmulti.Fanout(
			handlers...,
		),
	)
	l.Logger = logger

	return nil
}

func createAttr(args ...KeyVal) []slog.Attr {
	var attrs []slog.Attr
	for _, arg := range args {
		var value slog.Value
		if kvs, ok := arg.Value.([]KeyVal); ok {
			groupArgs := createAttr(kvs...)
			value = slog.GroupValue(groupArgs...)
		} else {
			value = slog.AnyValue(arg.Value)
		}
		attrs = append(attrs, slog.Attr{
			Key:   arg.Key,
			Value: value,
		})
	}

	return attrs
}

func createArgs(args ...KeyVal) []any {
	attrs := createAttr(args...)

	anyKeyVals := make([]any, 0, len(attrs)*2)
	for _, arg := range attrs {
		anyKeyVals = append(anyKeyVals, arg.Key, arg.Value)
	}
	return anyKeyVals
}

// With adds attributes to the logger
func (l *SlogLogger) With(args ...KeyVal) Logger {
	newLogger := l.Logger.With(createArgs(args...)...)
	return &SlogLogger{
		Logger: newLogger,
	}
}

// Debug logs a debug message
func (l *SlogLogger) Debug(ctx context.Context, msg string, args ...KeyVal) {
	l.Logger.DebugContext(ctx, msg, createArgs(args...)...)
}

// Info logs an info message
func (l *SlogLogger) Info(ctx context.Context, msg string, args ...KeyVal) {
	l.Logger.InfoContext(ctx, msg, createArgs(args...)...)
}

// Warn logs a warning message
func (l *SlogLogger) Warn(ctx context.Context, msg string, args ...KeyVal) {
	l.Logger.WarnContext(ctx, msg, createArgs(args...)...)
}

// Error logs an error message
func (l *SlogLogger) Error(ctx context.Context, msg string, args ...KeyVal) {
	l.Logger.ErrorContext(ctx, msg, createArgs(args...)...)
}

// Close closes the logger
func (l *SlogLogger) Close() error {
	for _, file := range l.openFIles {
		if err := file.Close(); err != nil {
			return fmt.Errorf("failed to close file: %w", err)
		}
	}
	return nil
}

var _ Logger = &SlogLogger{}
