package config

import "fmt"

// LoggingOutputType is the type of the logging output.
type LoggingOutputType string

const (
	// LoggingOutputTypeFile is the type of the file.
	LoggingOutputTypeFile LoggingOutputType = "file"
	// LoggingOutputTypeStdout is the type of the stdout.
	LoggingOutputTypeStdout LoggingOutputType = "stdout"
	// LoggingOutputTypeTCP is the type of the tcp.
	LoggingOutputTypeTCP LoggingOutputType = "tcp"
)

// LoggingOutputFormat is the format of the logging output.
type LoggingOutputFormat string

const (
	// LoggingOutputFormatJSON is the format of the json.
	LoggingOutputFormatJSON LoggingOutputFormat = "json"
	// LoggingOutputFormatText is the format of the text.
	LoggingOutputFormatText LoggingOutputFormat = "text"
)

// LoggingOutputLevel is the level of the logging output.
type LoggingOutputLevel string

const (
	// LoggingOutputLevelDebug is the level of the debug.
	LoggingOutputLevelDebug LoggingOutputLevel = "debug"
	// LoggingOutputLevelInfo is the level of the info.
	LoggingOutputLevelInfo LoggingOutputLevel = "info"
	// LoggingOutputLevelWarn is the level of the warn.
	LoggingOutputLevelWarn LoggingOutputLevel = "warn"
	// LoggingOutputLevelError is the level of the error.
	LoggingOutputLevelError LoggingOutputLevel = "error"
)

// LoggingOutputConfig represents the configuration for logging output
type LoggingOutputConfig struct {
	Type       *string   `mapstructure:"type"`
	Format     *string   `mapstructure:"format"`
	Level      *string   `mapstructure:"level"`
	Filename   *string   `mapstructure:"filename"`
	Address    *string   `mapstructure:"address"`
	EnabledEnv *[]string `mapstructure:"enabled_env"`
}

// ValidLoggingOutputConfig represents the valid configuration for logging output
type ValidLoggingOutputConfig struct {
	Type       LoggingOutputType
	Format     LoggingOutputFormat
	Level      LoggingOutputLevel
	Filename   string
	Address    string
	EnabledEnv struct {
		All    bool
		Values []string
	}
}

// LoggingConfig represents the configuration for logging
type LoggingConfig struct {
	Output []LoggingOutputConfig `mapstructure:"output"`
}

// ValidLoggingConfig represents the valid configuration for logging
type ValidLoggingConfig struct {
	Output []ValidLoggingOutputConfig
}

// Validate validates the logging configuration
func (l LoggingConfig) Validate() (ValidLoggingConfig, error) {
	var valid ValidLoggingConfig
	for i, output := range l.Output {
		var validOutput ValidLoggingOutputConfig
		if output.Type == nil {
			return ValidLoggingConfig{}, fmt.Errorf("output[%d].type: %w", i, ErrLoggingOutputTypeRequired)
		}
		if output.Format == nil {
			return ValidLoggingConfig{}, fmt.Errorf("output[%d].format: %w", i, ErrLoggingOutputFormatRequired)
		}
		if output.Level == nil {
			return ValidLoggingConfig{}, fmt.Errorf("output[%d].level: %w", i, ErrLoggingOutputLevelRequired)
		}
		switch *output.Type {
		case string(LoggingOutputTypeFile):
			validOutput.Type = LoggingOutputTypeFile
			if output.Filename == nil {
				return ValidLoggingConfig{}, fmt.Errorf("output[%d].filename: %w", i, ErrLoggingOutputFilenameRequired)
			}
			validOutput.Filename = *output.Filename
		case string(LoggingOutputTypeStdout):
			validOutput.Type = LoggingOutputTypeStdout
		case string(LoggingOutputTypeTCP):
			validOutput.Type = LoggingOutputTypeTCP
			if output.Address == nil {
				return ValidLoggingConfig{}, fmt.Errorf("output[%d].address: %w", i, ErrLoggingOutputAddressRequired)
			}
			validOutput.Address = *output.Address
		default:
			return ValidLoggingConfig{}, fmt.Errorf("output[%d].type: %w", i, ErrLoggingOutputTypeInvalid)
		}
		switch *output.Format {
		case string(LoggingOutputFormatJSON):
			validOutput.Format = LoggingOutputFormatJSON
		case string(LoggingOutputFormatText):
			validOutput.Format = LoggingOutputFormatText
		default:
			return ValidLoggingConfig{}, fmt.Errorf("output[%d].format: %w", i, ErrLoggingOutputFormatInvalid)
		}
		if output.Level == nil {
			validOutput.Level = LoggingOutputLevelInfo
		} else {
			switch *output.Level {
			case string(LoggingOutputLevelDebug):
				validOutput.Level = LoggingOutputLevelDebug
			case string(LoggingOutputLevelInfo):
				validOutput.Level = LoggingOutputLevelInfo
			case string(LoggingOutputLevelWarn):
				validOutput.Level = LoggingOutputLevelWarn
			case string(LoggingOutputLevelError):
				validOutput.Level = LoggingOutputLevelError
			default:
				return ValidLoggingConfig{}, fmt.Errorf("output[%d].level: %w", i, ErrLoggingOutputLevelInvalid)
			}
		}
		if output.EnabledEnv != nil {
			validOutput.EnabledEnv.Values = *output.EnabledEnv
		} else {
			validOutput.EnabledEnv.All = true
		}

		valid.Output = append(valid.Output, validOutput)
	}
	return valid, nil
}
