package runner

import (
	"fmt"
	"time"
)

// Kind represents the kind of runner
type Kind string

const (
	// RunnerKindStoreValue represents the store value runner
	RunnerKindStoreValue Kind = "StoreValue"
	// RunnerKindMemoryValue represents the memory store value runner
	RunnerKindMemoryValue Kind = "MemoryValue"
	// RunnerKindStoreImport represents the store import runner
	RunnerKindStoreImport Kind = "StoreImport"
	// RunnerKindOneExecute represents execute one request runner
	RunnerKindOneExecute Kind = "OneExecute"
	// RunnerKindMassExecute represents execute multiple requests runner
	RunnerKindMassExecute Kind = "MassExecute"
	// RunnerKindFlow represents the flow runner
	RunnerKindFlow Kind = "Flow"
	// RunnerKindSlaveConnect represents the slave connect runner
	RunnerKindSlaveConnect Kind = "SlaveConnect"
)

// Runner represents a runner
type Runner struct {
	Kind        *string        `yaml:"kind"`
	Sleep       Sleep          `yaml:"sleep"`
	StoreImport RunStoreImport `yaml:"store_import"`
}

// ValidRunner represents a valid runner
type ValidRunner struct {
	Kind        Kind
	Sleep       ValidRunnerSleep
	StoreImport ValidRunnerStoreImport
}

// Validate validates a runner
func (r Runner) Validate() (ValidRunner, error) {
	if r.Kind == nil {
		return ValidRunner{}, fmt.Errorf("kind is required")
	}
	var kind Kind
	switch Kind(*r.Kind) {
	case RunnerKindStoreValue,
		RunnerKindMemoryValue,
		RunnerKindStoreImport,
		RunnerKindOneExecute,
		RunnerKindMassExecute,
		RunnerKindFlow,
		RunnerKindSlaveConnect:
		kind = Kind(*r.Kind)
	default:
		return ValidRunner{}, fmt.Errorf("invalid kind value: %s", *r.Kind)
	}
	validSleep, err := r.Sleep.Validate()
	if err != nil {
		return ValidRunner{}, fmt.Errorf("failed to validate sleep: %w", err)
	}
	validStoreImport, err := r.StoreImport.Validate()
	if err != nil {
		return ValidRunner{}, fmt.Errorf("failed to validate store import: %w", err)
	}
	return ValidRunner{
		Kind:        kind,
		Sleep:       validSleep,
		StoreImport: validStoreImport,
	}, nil
}

// Sleep represents the sleep configuration for a runner
type Sleep struct {
	Enabled bool         `yaml:"enabled"`
	Values  []SleepValue `yaml:"values"`
}

// ValidRunnerSleep represents a valid runner sleep configuration
type ValidRunnerSleep struct {
	Enabled bool
	Values  []ValidRunnerSleepValue
}

// Validate validates a runnerSleep
func (r Sleep) Validate() (ValidRunnerSleep, error) {
	if !r.Enabled {
		return ValidRunnerSleep{}, nil
	}
	var values []ValidRunnerSleepValue
	for _, v := range r.Values {
		valid, err := v.Validate()
		if err != nil {
			return ValidRunnerSleep{}, fmt.Errorf("failed to validate sleep value: %w", err)
		}
		values = append(values, valid)
	}
	return ValidRunnerSleep{
		Enabled: r.Enabled,
		Values:  values,
	}, nil
}

// SleepValueAfter represents the after value for a runner sleep value
type SleepValueAfter string

const (
	// RunnerSleepValueAfterInit represents the init after value for a runner sleep value
	RunnerSleepValueAfterInit SleepValueAfter = "init"
	// RunnerSleepValueAfterExec represents the success after value for a runner sleep value
	RunnerSleepValueAfterExec SleepValueAfter = "exec"
	// RunnerSleepValueAfterFailedExec represents the failed after value for a runner sleep value
	RunnerSleepValueAfterFailedExec SleepValueAfter = "failedExec"
)

// SleepValue represents the sleep value for a runner
type SleepValue struct {
	Duration *string `yaml:"duration"`
	After    *string `yaml:"after"`
}

// ValidRunnerSleepValue represents a valid runner sleep value
type ValidRunnerSleepValue struct {
	Duration time.Duration
	After    SleepValueAfter
}

// Validate validates a runner
func (r SleepValue) Validate() (ValidRunnerSleepValue, error) {
	if r.Duration == nil {
		return ValidRunnerSleepValue{}, fmt.Errorf("duration is required")
	}
	if r.After == nil {
		return ValidRunnerSleepValue{}, fmt.Errorf("after is required")
	}
	d, err := time.ParseDuration(*r.Duration)
	if err != nil {
		return ValidRunnerSleepValue{}, fmt.Errorf("failed to parse duration: %w", err)
	}
	var after SleepValueAfter
	switch SleepValueAfter(*r.After) {
	case RunnerSleepValueAfterInit, RunnerSleepValueAfterExec, RunnerSleepValueAfterFailedExec:
		after = SleepValueAfter(*r.After)
	default:
		return ValidRunnerSleepValue{}, fmt.Errorf("invalid after value: %s", *r.After)
	}
	return ValidRunnerSleepValue{
		Duration: d,
		After:    after,
	}, nil
}

// RetrieveSleepValue retrieves the sleep value for a runner
func (r ValidRunner) RetrieveSleepValue(after SleepValueAfter) (time.Duration, bool) {
	for _, v := range r.Sleep.Values {
		if v.After == after {
			return v.Duration, true
		}
	}
	return time.Duration(0), false
}

// RunStoreImport represents the StoreImport runner
type RunStoreImport struct {
	Enabled bool              `yaml:"enabled"`
	Data    []StoreImportData `yaml:"data"`
}

// ValidRunnerStoreImport represents the valid RunnerStoreImport runner
type ValidRunnerStoreImport struct {
	Enabled bool
	Data    []ValidStoreImportData
}

// Validate validates the RunnerStoreImport
func (r RunStoreImport) Validate() (ValidRunnerStoreImport, error) {
	if !r.Enabled {
		return ValidRunnerStoreImport{}, nil
	}
	var validData []ValidStoreImportData
	for i, d := range r.Data {
		valid, err := d.Validate()
		if err != nil {
			return ValidRunnerStoreImport{}, fmt.Errorf("failed to validate data at index %d: %w", i, err)
		}
		validData = append(validData, valid)
	}
	return ValidRunnerStoreImport{
		Enabled: r.Enabled,
		Data:    validData,
	}, nil
}

// CredentialEncryptConfig is the configuration for the credential encrypt.
type CredentialEncryptConfig struct {
	Enabled   bool    `yaml:"enabled"`
	EncryptID *string `yaml:"encrypt_id"`
}

// ValidCredentialEncryptConfig represents the valid auth credential encrypt configuration
type ValidCredentialEncryptConfig struct {
	Enabled   bool
	EncryptID string
}

// Validate validates the credential encrypt configuration
func (c CredentialEncryptConfig) Validate() (ValidCredentialEncryptConfig, error) {
	if !c.Enabled {
		return ValidCredentialEncryptConfig{}, nil
	}
	if c.EncryptID == nil {
		return ValidCredentialEncryptConfig{}, fmt.Errorf("encrypt_id is required")
	}
	return ValidCredentialEncryptConfig{
		Enabled:   c.Enabled,
		EncryptID: *c.EncryptID,
	}, nil
}
