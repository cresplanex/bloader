package config

import "fmt"

// OutputType represents the type of the output service
type OutputType string

const (
	// OutputTypeLocal represents the Local output service
	OutputTypeLocal OutputType = "local"
)

// OutputFormat represents the format of the output service
type OutputFormat string

const (
	// OutputFormatCSV represents the CSV output service
	OutputFormatCSV OutputFormat = "csv"
)

// OutputRespectiveValueConfig represents the configuration for the output respective service value
type OutputRespectiveValueConfig struct {
	Env      *string `mapstructure:"env"`
	Type     *string `mapstructure:"type"`
	Format   *string `mapstructure:"format"`
	BasePath *string `mapstructure:"base_path"`
}

// ValidOutputRespectiveValueConfig represents the configuration for the output respective service value
type ValidOutputRespectiveValueConfig struct {
	Env      string
	Type     OutputType
	Format   OutputFormat
	BasePath string
}

// Validate validates the output respective value configuration
func (c OutputRespectiveValueConfig) Validate() (ValidOutputRespectiveValueConfig, error) {
	var valid ValidOutputRespectiveValueConfig
	if c.Env == nil {
		return ValidOutputRespectiveValueConfig{}, ErrOutputValueEnvRequired
	}
	valid.Env = *c.Env

	if c.Type == nil {
		return ValidOutputRespectiveValueConfig{}, ErrOutputValueTypeRequired
	}
	switch OutputType(*c.Type) {
	case OutputTypeLocal:
		valid.Type = OutputTypeLocal
		if c.Format == nil {
			return ValidOutputRespectiveValueConfig{}, ErrOutputValueFormatRequired
		}
		switch OutputFormat(*c.Format) {
		case OutputFormatCSV:
			valid.Format = OutputFormatCSV
		default:
			return ValidOutputRespectiveValueConfig{}, ErrOutputValueFormatInvalid
		}
		if c.BasePath == nil {
			return ValidOutputRespectiveValueConfig{}, ErrOutputValueBasePathRequired
		}
		valid.BasePath = *c.BasePath
	default:
		return ValidOutputRespectiveValueConfig{}, ErrOutputValueTypeInvalid
	}

	return valid, nil
}

// OutputRespectiveConfig is a struct that represents the output configuration
type OutputRespectiveConfig struct {
	ID     *string                       `mapstructure:"id"`
	Values []OutputRespectiveValueConfig `mapstructure:"values"`
}

// ValidOutputRespectiveConfig represents the configuration for the output respective service
type ValidOutputRespectiveConfig struct {
	ID     string
	Values []ValidOutputRespectiveValueConfig
}

// OutputConfig represents the configuration for the output service
type OutputConfig []OutputRespectiveConfig

// ValidOutputConfig represents the configuration for the output service
type ValidOutputConfig []ValidOutputRespectiveConfig

// Validate validates the output configuration
func (c OutputConfig) Validate() (ValidOutputConfig, error) {
	var valid ValidOutputConfig
	idSet := make(map[string]struct{})
	for i, output := range c {
		var validRespective ValidOutputRespectiveConfig
		if output.ID == nil {
			return ValidOutputConfig{}, ErrOutputValueIDRequired
		}
		if _, ok := idSet[*output.ID]; ok {
			return ValidOutputConfig{}, ErrOutputValueIDDuplicate
		}
		idSet[*output.ID] = struct{}{}
		validRespective.ID = *output.ID

		var validValues []ValidOutputRespectiveValueConfig
		for j, value := range output.Values {
			validValue, err := value.Validate()
			if err != nil {
				return ValidOutputConfig{}, fmt.Errorf("output[%d].values[%d]: %w", i, j, err)
			}
			validValues = append(validValues, validValue)
		}
		validRespective.Values = validValues
		valid = append(valid, validRespective)
	}
	return valid, nil
}
