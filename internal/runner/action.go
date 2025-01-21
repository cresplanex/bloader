package runner

import (
	"context"
	"fmt"
)

// ActionType represents the action
type ActionType string

const (
	// ActionTypeTermWithErr represents the term with error
	ActionTypeTermWithErr ActionType = "sys:term_with_err"
	// ActionTypeTermWithoutErr represents the term without error
	ActionTypeTermWithoutErr ActionType = "sys:term_without_err"
)

// Actions represents the actions
type Actions []Action

// ValidActions represents the valid actions
type ValidActions []ValidAction

// Validate validates the actions
func (a Actions) Validate() (ValidActions, error) {
	var valids ValidActions
	idSet := make(map[string]struct{})
	for i, action := range a {
		var valid ValidAction
		if action.ID == nil {
			return ValidActions{}, fmt.Errorf("auth[%d].id is required", i)
		}
		if _, ok := idSet[*action.ID]; ok {
			return nil, fmt.Errorf("duplicate id: %s", *action.ID)
		}
		idSet[*action.ID] = struct{}{}
		valid.ID = *action.ID
		if err := action.Validate(&valid); err != nil {
			return nil, fmt.Errorf("failed to validate action at index %d: %w", i, err)
		}
		valids = append(valids, valid)
	}
	return valids, nil
}

// Action represents the action
type Action struct {
	ID   *string `yaml:"id"`
	Type *string `yaml:"type"`
	On   []ActionOn
}

// ValidAction represents the valid action
type ValidAction struct {
	ID   string
	Type ActionType
	On   []ValidActionOn
}

// Validate validates the action
func (a Action) Validate(valid *ValidAction) error {
	if a.ID == nil {
		return fmt.Errorf("id is required")
	}
	valid.ID = *a.ID
	if a.Type == nil {
		return fmt.Errorf("type is required")
	}
	valid.Type = ActionType(*a.Type)
	if a.On == nil {
		return fmt.Errorf("on is required")
	}
	var validOns []ValidActionOn
	for i, on := range a.On {
		validOn, err := on.Validate()
		if err != nil {
			return fmt.Errorf("failed to validate on at index %d: %w", i, err)
		}
		validOns = append(validOns, validOn)
	}
	valid.On = validOns
	return nil
}

// ActionOn represents the action on
type ActionOn struct {
	Flow  *string `yaml:"flow"`
	Event *string `yaml:"event"`
}

// ValidActionOn represents the valid action on
type ValidActionOn struct {
	Flow  string
	Event Event
}

// Validate validates the action
func (a ActionOn) Validate() (ValidActionOn, error) {
	var valid ValidActionOn
	if a.Flow == nil {
		return ValidActionOn{}, fmt.Errorf("flow is required")
	}
	valid.Flow = *a.Flow
	if a.Event == nil {
		return ValidActionOn{}, fmt.Errorf("event is required")
	}
	valid.Event = Event(*a.Event)
	return valid, nil
}

// ActionCaster represents the action caster
type ActionCaster map[ActionType]chan struct{}

// NewActionCasterFromRunnerKind creates a new action caster from the runner kind
func NewActionCasterFromRunnerKind(kind Kind) (ActionCaster, error) {
	var actionTypes []ActionType
	switch kind {
	case RunnerKindStoreValue:
		actionTypes = StoreValueRunnerActionsList
	case RunnerKindMemoryValue:
		actionTypes = MemoryValueRunnerActionsList
	case RunnerKindStoreImport:
		actionTypes = StoreImportRunnerActionsList
	case RunnerKindOneExecute:
		actionTypes = OneExecuteRunnerActionsList
	case RunnerKindMassExecute:
		actionTypes = MassExecuteRunnerActionsList
	case RunnerKindSlaveConnect:
		actionTypes = SlaveConnectRunnerActionsList
	case RunnerKindFlow:
		actionTypes = FlowRunnerActionsList
	default:
		return nil, fmt.Errorf("unsupported runner kind: %s", kind)
	}

	caster := make(ActionCaster)
	for _, actionType := range actionTypes {
		caster[actionType] = make(chan struct{})
	}

	return nil, fmt.Errorf("unsupported runner kind: %s", kind)
}

// FindChannel finds the channel
func (c ActionCaster) FindChannel(actionType ActionType) (chan struct{}, bool) {
	ch, ok := c[actionType]
	return ch, ok
}

// Send sends the action caster
func (c ActionCaster) Send(ctx context.Context, actionType ActionType) {
	ch, ok := c.FindChannel(actionType)
	if ok {
		select {
		case ch <- struct{}{}:
		case <-ctx.Done():
		}
	}
}

// Close closes the action caster
func (c ActionCaster) Close() {
	for _, ch := range c {
		close(ch)
	}
}
