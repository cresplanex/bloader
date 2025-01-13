package config

import "fmt"

// EncryptRespectiveConfig is the configuration for the encrypt.
type EncryptRespectiveConfig struct {
	ID    *string             `mapstructure:"id"`
	Type  *string             `mapstructure:"type"`
	Key   *string             `mapstructure:"key"`
	Store *StoreSpecifyConfig `mapstructure:"store"`
}

// ValidEncryptRespectiveConfig represents the valid encrypt configuration
type ValidEncryptRespectiveConfig struct {
	ID    string
	Type  EncryptType
	Key   []byte
	Store ValidStoreSpecifyConfig
}

// EncryptType is the type of the encrypt.
type EncryptType string

const (
	// EncryptTypeStaticCBC is the type of the static cbc.
	EncryptTypeStaticCBC EncryptType = "staticCBC"
	// EncryptTypeStaticCFB is the type of the static cfb.
	EncryptTypeStaticCFB EncryptType = "staticCFB"
	// EncryptTypeStaticCTR is the type of the static ctr.
	EncryptTypeStaticCTR EncryptType = "staticCTR"
	// EncryptTypeDynamicCBC is the type of the dynamic cbc.
	EncryptTypeDynamicCBC EncryptType = "dynamicCBC"
	// EncryptTypeDynamicCFB is the type of the dynamic cfb.
	EncryptTypeDynamicCFB EncryptType = "dynamicCFB"
	// EncryptTypeDynamicCTR is the type of the dynamic ctr.
	EncryptTypeDynamicCTR EncryptType = "dynamicCTR"
)

// EncryptConfig is the configuration for the encrypt.
type EncryptConfig []EncryptRespectiveConfig

// ValidEncryptConfig represents the valid encrypt configuration
type ValidEncryptConfig []ValidEncryptRespectiveConfig

// Validate validates the encrypt configuration.
func (c EncryptConfig) Validate() (ValidEncryptConfig, error) {
	var valid ValidEncryptConfig
	idSet := make(map[string]struct{})
	for i, ec := range c {
		var validRespective ValidEncryptRespectiveConfig
		if ec.ID == nil {
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].id: %w", i, ErrEncryptIDRequired)
		}
		if _, ok := idSet[*ec.ID]; ok {
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].id: %w", i, ErrEncryptIDDuplicate)
		}
		idSet[*ec.ID] = struct{}{}
		validRespective.ID = *ec.ID
		if ec.Type == nil {
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].type: %w", i, ErrEncryptTypeRequired)
		}
		switch EncryptType(*ec.Type) {
		case EncryptTypeStaticCBC, EncryptTypeStaticCFB, EncryptTypeStaticCTR:
			validRespective.Type = EncryptType(*ec.Type)
			if ec.Key == nil {
				return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].key: %w", i, ErrEncryptKeyRequired)
			}
			validRespective.Key = []byte(*ec.Key)
			if len(*ec.Key) != 16 && len(*ec.Key) != 24 && len(*ec.Key) != 32 {
				return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].key: %w", i, ErrEncryptRSAKeySizeInvalid)
			}
		case EncryptTypeDynamicCBC, EncryptTypeDynamicCFB, EncryptTypeDynamicCTR:
			validRespective.Type = EncryptType(*ec.Type)
			if ec.Store == nil {
				return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].store: %w", i, ErrEncryptStoreRequired)
			}
			validStore, err := ec.Store.Validate()
			if err != nil {
				return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].store: %w", i, err)
			}
			validRespective.Store = validStore
		default:
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].type: %w", i, ErrEncryptTypeInvalid)
		}
		valid = append(valid, validRespective)
	}
	return valid, nil
}

// ValidateOnSlave validates the encrypt configuration on the slave.
func (c EncryptConfig) ValidateOnSlave() (ValidEncryptConfig, error) {
	var valid ValidEncryptConfig
	idSet := make(map[string]struct{})
	for i, ec := range c {
		var validRespective ValidEncryptRespectiveConfig
		if ec.ID == nil {
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].id: %w", i, ErrEncryptIDRequired)
		}
		if _, ok := idSet[*ec.ID]; ok {
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].id: %w", i, ErrEncryptIDDuplicate)
		}
		idSet[*ec.ID] = struct{}{}
		validRespective.ID = *ec.ID
		if ec.Type == nil {
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].type: %w", i, ErrEncryptTypeRequired)
		}
		switch EncryptType(*ec.Type) {
		case EncryptTypeStaticCBC, EncryptTypeStaticCFB, EncryptTypeStaticCTR:
			validRespective.Type = EncryptType(*ec.Type)
			if ec.Key == nil {
				return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].key: %w", i, ErrEncryptKeyRequired)
			}
			validRespective.Key = []byte(*ec.Key)
			if len(*ec.Key) != 16 && len(*ec.Key) != 24 && len(*ec.Key) != 32 {
				return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].key: %w", i, ErrEncryptRSAKeySizeInvalid)
			}
		case EncryptTypeDynamicCBC, EncryptTypeDynamicCFB, EncryptTypeDynamicCTR:
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].type: %w", i, ErrEncryptTypeUnsupportedOnSlave)
		default:
			return ValidEncryptConfig{}, fmt.Errorf("encrypt[%d].type: %w", i, ErrEncryptTypeInvalid)
		}
		valid = append(valid, validRespective)
	}
	return valid, nil
}
