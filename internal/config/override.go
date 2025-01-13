package config

import "fmt"

// OverrideType represents the type of the override service
type OverrideType string

const (
	// OverrideTypeStatic represents the Static override service
	OverrideTypeStatic OverrideType = "static"
	// OverrideTypeFile represents the File override service
	OverrideTypeFile OverrideType = "file"
)

// OverrideRespectiveVarConfig represents the configuration for the override respective service var
type OverrideRespectiveVarConfig struct {
	Key   *string `mapstructure:"key"`
	Value *string `mapstructure:"value"`
}

// ValidOverrideRespectiveVarConfig represents the configuration for the override respective service var
type ValidOverrideRespectiveVarConfig struct {
	Key   string
	Value string
}

// Validate validates the override respective var configuration
func (c OverrideRespectiveVarConfig) Validate() (ValidOverrideRespectiveVarConfig, error) {
	var valid ValidOverrideRespectiveVarConfig
	if c.Key == nil {
		return ValidOverrideRespectiveVarConfig{}, ErrOverrideVarKeyRequired
	}
	valid.Key = *c.Key

	if c.Value == nil {
		return ValidOverrideRespectiveVarConfig{}, ErrOverrideVarValueRequired
	}
	valid.Value = *c.Value

	return valid, nil
}

// OverrideFileTypes represents the file types for the override respective service
type OverrideFileTypes string

const (
	// OverrideFileTypesYAML represents the YAML file type for the override respective service
	OverrideFileTypesYAML OverrideFileTypes = "yaml"
)

// OverrideRespectiveConfig represents the configuration for the override respective service
type OverrideRespectiveConfig struct {
	Type       *string                       `mapstructure:"type"`
	FileType   *string                       `mapstructure:"file_type"`
	Path       *string                       `mapstructure:"path"`
	Partial    bool                          `mapstructure:"partial"`
	Vars       []OverrideRespectiveVarConfig `mapstructure:"vars"`
	Store      *StoreSpecifyConfig           `mapstructure:"store"`
	Key        *string                       `mapstructure:"key"`
	Value      *string                       `mapstructure:"value"`
	EnabledEnv *[]string                     `mapstructure:"enabled_env"`
}

// ValidOverrideRespectiveConfig represents the configuration for the override respective service
type ValidOverrideRespectiveConfig struct {
	Type       OverrideType
	FileType   OverrideFileTypes
	Path       string
	Partial    bool
	Vars       []ValidOverrideRespectiveVarConfig
	Key        string
	Value      string
	EnabledEnv struct {
		All    bool
		Values []string
	}
}

// OverrideConfig represents the configuration for the override service
type OverrideConfig []OverrideRespectiveConfig

// ValidOverrideConfig represents the configuration for the override service
type ValidOverrideConfig []ValidOverrideRespectiveConfig

// Validate validates the override configuration
func (c OverrideConfig) Validate() (ValidOverrideConfig, error) {
	var valid ValidOverrideConfig
	for i, override := range c {
		var validOverride ValidOverrideRespectiveConfig
		if override.Type == nil {
			return nil, fmt.Errorf("override[%d]: %w", i, ErrOverrideTypeRequired)
		}
		switch OverrideType(*override.Type) {
		case OverrideTypeStatic:
			validOverride.Type = OverrideTypeStatic
			if override.Key == nil {
				return nil, fmt.Errorf("override[%d]: %w", i, ErrOverrideKeyRequired)
			}
			validOverride.Key = *override.Key
			if override.Value == nil {
				return nil, fmt.Errorf("override[%d]: %w", i, ErrOverrideValueRequired)
			}
			validOverride.Value = *override.Value
		case OverrideTypeFile:
			validOverride.Type = OverrideTypeFile
			if override.FileType == nil {
				return nil, fmt.Errorf("override[%d]: %w", i, ErrOverrideFileTypeRequired)
			}
			switch OverrideFileTypes(*override.FileType) {
			case OverrideFileTypesYAML:
				validOverride.FileType = OverrideFileTypesYAML
			default:
				return nil, fmt.Errorf("override[%d]: %w", i, ErrOverrideFileTypeInvalid)
			}
			if override.Path == nil {
				return nil, fmt.Errorf("override[%d]: %w", i, ErrOverridePathRequired)
			}
			validOverride.Path = *override.Path
			validOverride.Partial = override.Partial
			if override.Partial {
				for j, varConfig := range override.Vars {
					validVarConfig, err := varConfig.Validate()
					if err != nil {
						return nil, fmt.Errorf("override[%d].vars[%d]: %w", i, j, err)
					}
					validOverride.Vars = append(validOverride.Vars, validVarConfig)
				}
			}
		default:
			return nil, ErrOverrideTypeInvalid
		}
		if override.EnabledEnv != nil {
			validOverride.EnabledEnv.Values = *override.EnabledEnv
		} else {
			validOverride.EnabledEnv.All = true
		}

		valid = append(valid, validOverride)
	}

	return valid, nil
}
