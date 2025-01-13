package matcher

import (
	"fmt"

	"github.com/jmespath/go-jmespath"
)

// DataExtractorOnNilType represents the on nil type for the data extractor
type DataExtractorOnNilType string

const (
	// DataExtractorOnNilTypeEmpty represents the empty on nil type for the data extractor
	DataExtractorOnNilTypeEmpty DataExtractorOnNilType = "empty"
	// DataExtractorOnNilTypeNull represents the null on nil type for the data extractor
	DataExtractorOnNilTypeNull DataExtractorOnNilType = "null"
	// DataExtractorOnNilTypeError represents the error on nil type for the data extractor
	DataExtractorOnNilTypeError DataExtractorOnNilType = "error"

	// DefaultDataExtractorOnNilType represents the default on nil type for the data extractor
	DefaultDataExtractorOnNilType DataExtractorOnNilType = DataExtractorOnNilTypeNull
)

// DataExtractorType represents the type for the data extractor
type DataExtractorType string

const (
	// DataExtractorTypeJMESPath represents the JMESPath type for the data extractor
	DataExtractorTypeJMESPath DataExtractorType = "jmesPath"
)

// DataExtractor represents the data extractor for the OneExec runner
type DataExtractor struct {
	Type     *string `yaml:"type"`
	JMESPath *string `yaml:"jmes_path"`
	OnNil    *string `yaml:"on_nil"`
}

// ValidDataExtractor represents the valid data extractor for the OneExec runner
type ValidDataExtractor struct {
	Type     DataExtractorType
	JMESPath *jmespath.JMESPath
	OnNil    DataExtractorOnNilType
}

// Validate validates the data extractor
func (d DataExtractor) Validate() (ValidDataExtractor, error) {
	if d.Type == nil {
		return ValidDataExtractor{}, fmt.Errorf("type is required")
	}
	var valid ValidDataExtractor
	switch DataExtractorType(*d.Type) {
	case DataExtractorTypeJMESPath:
		valid.Type = DataExtractorType(*d.Type)
		if d.JMESPath == nil {
			return ValidDataExtractor{}, fmt.Errorf("jmesPath is required")
		}
		jPath, err := jmespath.Compile(*d.JMESPath)
		if err != nil {
			return ValidDataExtractor{}, fmt.Errorf("failed to compile jmesPath: %w", err)
		}
		valid.JMESPath = jPath
		if d.OnNil == nil {
			valid.OnNil = DefaultDataExtractorOnNilType
		} else {
			switch DataExtractorOnNilType(*d.OnNil) {
			case DataExtractorOnNilTypeEmpty, DataExtractorOnNilTypeNull, DataExtractorOnNilTypeError:
				valid.OnNil = DataExtractorOnNilType(*d.OnNil)
			default:
				valid.OnNil = DefaultDataExtractorOnNilType
			}
		}
	default:
		return ValidDataExtractor{}, fmt.Errorf("invalid type value: %s", *d.Type)
	}
	return valid, nil
}

// Extract extracts the data from the response
func (d ValidDataExtractor) Extract(data any) (any, error) {
	switch d.Type {
	case DataExtractorTypeJMESPath:
		result, err := d.JMESPath.Search(data)
		if err != nil {
			return nil, fmt.Errorf("failed to search jmesPath: %w", err)
		}
		if result == nil {
			switch d.OnNil {
			case DataExtractorOnNilTypeEmpty:
				return "", nil
			case DataExtractorOnNilTypeNull:
				return nil, nil
			case DataExtractorOnNilTypeError:
				return nil, fmt.Errorf("nil value")
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported data extractor type: %s", d.Type)
	}
}
