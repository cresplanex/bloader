package config

import "time"

// ClockConfig represents the configuration for the clock
type ClockConfig struct {
	Format *string        `mapstructure:"format"`
	Fake   FakeTimeConfig `mapstructure:"fake"`
}

// ValidClockConfig represents the valid clock configuration
type ValidClockConfig struct {
	Format string
	Fake   ValidFakeTimeConfig
}

// DefaultClockFormat is the default clock format
const DefaultClockFormat = "2006-01-02T15:04:05Z"

// FakeTimeConfig represents the configuration for the fake time
type FakeTimeConfig struct {
	Enabled bool    `mapstructure:"enabled"`
	Time    *string `mapstructure:"time"`
}

// ValidFakeTimeConfig represents the valid fake time configuration
type ValidFakeTimeConfig struct {
	Enabled bool
	Time    time.Time
}

// Validate validates the clock configuration.
func (c ClockConfig) Validate() (ValidClockConfig, error) {
	var valid ValidClockConfig
	if c.Format == nil {
		valid.Format = DefaultClockFormat
	} else {
		valid.Format = *c.Format
	}
	if _, err := time.Parse(valid.Format, "2006-01-02T15:04:05Z"); err != nil {
		return ValidClockConfig{}, ErrClockFormatInvalid
	}
	valid.Fake.Enabled = c.Fake.Enabled
	if c.Fake.Enabled {
		if c.Fake.Time == nil {
			return ValidClockConfig{}, ErrClockFakeTimeRequired
		}
		fakeTime, err := time.Parse(valid.Format, *c.Fake.Time)
		if err != nil {
			return ValidClockConfig{}, ErrClockFakeTimeInvalid
		}
		valid.Fake.Time = fakeTime
	}
	return valid, nil
}
