package runner

import "fmt"

// ActionType represents the action
type ActionType string

const (
	// TermWithErr represents the term with error
	TermWithErr ActionType = "sys:term_with_err"
	// TermWithoutErr represents the term without error
	TermWithoutErr ActionType = "sys:term_without_err"
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
