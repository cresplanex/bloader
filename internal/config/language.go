package config

// LanguageType is the type of the language.
type LanguageType string

const (
	// LanguageTypeEnglish is the type of the english.
	LanguageTypeEnglish LanguageType = "en"
	// LanguageTypeJapanese is the type of the japanese.
	LanguageTypeJapanese LanguageType = "ja"
)

// LanguageConfig represents the language configuration
type LanguageConfig struct {
	Default *string `mapstructure:"default"`
}

// ValidLanguageConfig represents the valid language configuration
type ValidLanguageConfig struct {
	Default LanguageType
}

// Validate validates the language configuration
func (l LanguageConfig) Validate() (ValidLanguageConfig, error) {
	var valid ValidLanguageConfig
	if l.Default == nil {
		return ValidLanguageConfig{}, ErrLanguageDefaultRequired
	}
	switch *l.Default {
	case string(LanguageTypeEnglish):
		valid.Default = LanguageTypeEnglish
	case string(LanguageTypeJapanese):
		valid.Default = LanguageTypeJapanese
	default:
		return ValidLanguageConfig{}, ErrLanguageDefaultInvalid
	}

	return valid, nil
}
