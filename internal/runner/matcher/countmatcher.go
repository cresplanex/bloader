package matcher

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/cresplanex/bloader/internal/logger"
)

// CountOperator represents the status code operator
type CountOperator string

const (
	// CountOperatorNone represents the none operator
	CountOperatorNone CountOperator = "none"
	// CountOperatorEqual represents the equal operator
	CountOperatorEqual CountOperator = "eq"
	// CountOperatorNotEqual represents the not equal operator
	CountOperatorNotEqual CountOperator = "ne"
	// CountOperatorLessThan represents the less than operator
	CountOperatorLessThan CountOperator = "lt"
	// CountOperatorLessEqual represents the less equal operator
	CountOperatorLessEqual CountOperator = "le"
	// CountOperatorGreaterThan represents the greater than operator
	CountOperatorGreaterThan CountOperator = "gt"
	// CountOperatorGreaterEqual represents the greater equal operator
	CountOperatorGreaterEqual CountOperator = "ge"
	// CountOperatorIn represents the in operator
	CountOperatorIn CountOperator = "in"
	// CountOperatorNotIn represents the not in operator
	CountOperatorNotIn CountOperator = "nin"
	// CountOperatorBetween represents the between operator
	CountOperatorBetween CountOperator = "between"
	// CountOperatorNotBetween represents the not between operator
	CountOperatorNotBetween CountOperator = "notBetween"
	// CountOperatorMod represents the mod operator
	CountOperatorMod CountOperator = "mod"
	// CountOperatorNotMod represents the not mod operator
	CountOperatorNotMod CountOperator = "notMod"
	// CountOperatorRegex represents the regex operator
	CountOperatorRegex CountOperator = "regex"
)

// CountCondition represents the status code condition
type CountCondition struct {
	ID    *string `yaml:"id"`
	Op    *string `yaml:"op"`
	Value *any    `yaml:"value"`
}

type countBetweenVal struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// CountConditionMatcher represents the status code matcher
type CountConditionMatcher func(count int) bool

// MatcherGenerate generates the status code matcher
func (scc CountCondition) MatcherGenerate(ctx context.Context, log logger.Logger) (CountConditionMatcher, error) {
	if scc.ID == nil {
		return nil, fmt.Errorf("id is required")
	}
	if scc.Op == nil {
		return nil, fmt.Errorf("operator is required")
	}
	switch (CountOperator)(*scc.Op) {
	case CountOperatorNone:
		return func(_ int) bool {
			return false
		}, nil
	case CountOperatorEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count == countInt
		}, nil
	case CountOperatorNotEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count != countInt
		}, nil
	case CountOperatorLessThan:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count < countInt
		}, nil
	case CountOperatorLessEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count <= countInt
		}, nil
	case CountOperatorGreaterThan:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count > countInt
		}, nil
	case CountOperatorGreaterEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count >= countInt
		}, nil
	case CountOperatorIn:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		rawCodes, ok := (*scc.Value).([]any)
		if !ok {
			return nil, fmt.Errorf("value must be []int")
		}
		codes := make([]int, 0, len(rawCodes))
		for _, rawCode := range rawCodes {
			code, ok := rawCode.(int)
			if !ok {
				return nil, fmt.Errorf("value must be []int")
			}
			codes = append(codes, code)
		}
		return func(count int) bool {
			for _, v := range codes {
				if count == v {
					return true
				}
			}
			return false
		}, nil
	case CountOperatorNotIn:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		rawCodes, ok := (*scc.Value).([]any)
		if !ok {
			return nil, fmt.Errorf("value must be []int")
		}
		codes := make([]int, 0, len(rawCodes))
		for _, rawCode := range rawCodes {
			code, ok := rawCode.(int)
			if !ok {
				return nil, fmt.Errorf("value must be []int")
			}
			codes = append(codes, code)
		}
		return func(count int) bool {
			for _, v := range codes {
				if count == v {
					return false
				}
			}
			return true
		}, nil
	case CountOperatorBetween:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countVals, ok := (*scc.Value).(countBetweenVal)
		if !ok {
			return nil, fmt.Errorf("value must be countBetweenVal")
		}
		return func(count int) bool {
			return count >= countVals.Min && count <= countVals.Max
		}, nil
	case CountOperatorNotBetween:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countVals, ok := (*scc.Value).(countBetweenVal)
		if !ok {
			return nil, fmt.Errorf("value must be countBetweenVal")
		}
		return func(count int) bool {
			return count < countVals.Min || count > countVals.Max
		}, nil
	case CountOperatorMod:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count%countInt == 0
		}, nil
	case CountOperatorNotMod:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		countInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(count int) bool {
			return count%countInt != 0
		}, nil
	case CountOperatorRegex:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		strV, ok := (*scc.Value).(string)
		if !ok {
			return nil, fmt.Errorf("value must be string")
		}
		preRequireRegex := regexp.MustCompile(strV)
		return func(count int) bool {
			return preRequireRegex.MatchString(strconv.Itoa(count))
		}, nil
	default:
		log.Error(ctx, "unknown operator",
			logger.Value("operator", scc.Op))
		return nil, fmt.Errorf("unknown operator: %s", *scc.Op)
	}
}

// CountConditions represents the status code conditions
type CountConditions []CountCondition

// CountConditionsMatcher represents the status code conditions matcher
type CountConditionsMatcher func(count int) (string, bool)

// MatcherGenerate generates the status code conditions matcher
func (sccs CountConditions) MatcherGenerate(ctx context.Context, log logger.Logger) (CountConditionsMatcher, error) {
	matchers := make([]CountConditionMatcher, 0, len(sccs))
	for _, scc := range sccs {
		matcher, err := scc.MatcherGenerate(ctx, log)
		if err != nil {
			return nil, fmt.Errorf("failed to generate matcher: %w", err)
		}
		matchers = append(matchers, matcher)
	}
	return func(count int) (string, bool) {
		for i, matcher := range matchers {
			if matcher(count) {
				if sccs[i].ID != nil {
					return *sccs[i].ID, true
				}
				return "", true
			}
		}
		return "", false
	}, nil
}
