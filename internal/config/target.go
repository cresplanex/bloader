package config

import "fmt"

// TargetType represents the type of the target service
type TargetType string

const (
	// TargetTypeHTTP represents the HTTP target service
	TargetTypeHTTP TargetType = "http"
)

// TargetRespectiveValueConfig represents the configuration for the target respective service value
type TargetRespectiveValueConfig struct {
	Env *string `mapstructure:"env"`
	URL *string `mapstructure:"url"`
}

// ValidTargetRespectiveValueConfig represents the configuration for the target respective service value
type ValidTargetRespectiveValueConfig struct {
	Env string
	URL string
}

// Validate validates the target respective value configuration
func (c TargetRespectiveValueConfig) Validate() (ValidTargetRespectiveValueConfig, error) {
	var valid ValidTargetRespectiveValueConfig
	if c.Env == nil {
		return ValidTargetRespectiveValueConfig{}, ErrTargetValueEnvRequired
	}
	valid.Env = *c.Env

	if c.URL == nil {
		return ValidTargetRespectiveValueConfig{}, ErrTargetValueURLRequired
	}
	valid.URL = *c.URL

	return valid, nil
}

// TargetRespectiveConfig represents the configuration for the target respective service
type TargetRespectiveConfig struct {
	ID     *string                       `mapstructure:"id"`
	Type   *string                       `mapstructure:"type"`
	Values []TargetRespectiveValueConfig `mapstructure:"values"`
}

// ValidTargetRespectiveConfig represents the configuration for the target respective service
type ValidTargetRespectiveConfig struct {
	ID     string
	Type   TargetType
	Values []ValidTargetRespectiveValueConfig
}

// TargetConfig represents the configuration for the target service
type TargetConfig []TargetRespectiveConfig

// ValidTargetConfig represents the configuration for the target service
type ValidTargetConfig []ValidTargetRespectiveConfig

// Validate validates the target configuration
func (c TargetConfig) Validate() (ValidTargetConfig, error) {
	var valid ValidTargetConfig
	idSet := make(map[string]struct{})
	for i, target := range c {
		var validRespective ValidTargetRespectiveConfig
		if target.ID == nil {
			return ValidTargetConfig{}, fmt.Errorf("target[%d].id: %w", i, ErrTargetIDRequired)
		}
		if _, ok := idSet[*target.ID]; ok {
			return ValidTargetConfig{}, fmt.Errorf("target[%d].id: %w", i, ErrTargetIDDuplicate)
		}
		idSet[*target.ID] = struct{}{}
		validRespective.ID = *target.ID

		if target.Type == nil {
			return ValidTargetConfig{}, fmt.Errorf("target[%d].type: %w", i, ErrTargetTypeRequired)
		}
		switch *target.Type {
		case string(TargetTypeHTTP):
			validRespective.Type = TargetTypeHTTP
		default:
			return ValidTargetConfig{}, fmt.Errorf("target[%d].type: %w", i, ErrTargetTypeInvalid)
		}
		var validValues []ValidTargetRespectiveValueConfig
		for j, value := range target.Values {
			validValue, err := value.Validate()
			if err != nil {
				return ValidTargetConfig{}, fmt.Errorf("target[%d].values[%d]: %w", i, j, err)
			}
			validValues = append(validValues, validValue)
		}
		validRespective.Values = validValues
		valid = append(valid, validRespective)
	}

	return valid, nil
}
