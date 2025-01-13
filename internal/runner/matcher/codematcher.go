package matcher

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/cresplanex/bloader/internal/logger"
)

// StatusCodeOperator represents the status code operator
type StatusCodeOperator string

const (
	// StatusCodeOperatorNone represents the none operator
	StatusCodeOperatorNone StatusCodeOperator = "none"
	// StatusCodeOperatorEqual represents the equal operator
	StatusCodeOperatorEqual StatusCodeOperator = "eq"
	// StatusCodeOperatorNotEqual represents the not equal operator
	StatusCodeOperatorNotEqual StatusCodeOperator = "ne"
	// StatusCodeOperatorLessThan represents the less than operator
	StatusCodeOperatorLessThan StatusCodeOperator = "lt"
	// StatusCodeOperatorLessEqual represents the less equal operator
	StatusCodeOperatorLessEqual StatusCodeOperator = "le"
	// StatusCodeOperatorGreaterThan represents the greater than operator
	StatusCodeOperatorGreaterThan StatusCodeOperator = "gt"
	// StatusCodeOperatorGreaterEqual represents the greater equal operator
	StatusCodeOperatorGreaterEqual StatusCodeOperator = "ge"
	// StatusCodeOperatorIn represents the in operator
	StatusCodeOperatorIn StatusCodeOperator = "in"
	// StatusCodeOperatorNotIn represents the not in operator
	StatusCodeOperatorNotIn StatusCodeOperator = "nin"
	// StatusCodeOperatorBetween represents the between operator
	StatusCodeOperatorBetween StatusCodeOperator = "between"
	// StatusCodeOperatorNotBetween represents the not between operator
	StatusCodeOperatorNotBetween StatusCodeOperator = "notBetween"
	// StatusCodeOperatorRegex represents the regex operator
	StatusCodeOperatorRegex StatusCodeOperator = "regex"
)

// StatusCodeCondition represents the status code condition
type StatusCodeCondition struct {
	ID    *string `yaml:"id"`
	Op    *string `yaml:"op"`
	Value *any    `yaml:"value"`
}

type statusCodeBetweenVal struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// StatusCodeConditionMatcher represents the status code matcher
type StatusCodeConditionMatcher func(statusCode int) bool

// MatcherGenerate generates the status code matcher
func (scc StatusCodeCondition) MatcherGenerate(
	ctx context.Context,
	log logger.Logger,
) (StatusCodeConditionMatcher, error) {
	if scc.ID == nil {
		return nil, fmt.Errorf("id is required")
	}
	if scc.Op == nil {
		return nil, fmt.Errorf("operator is required")
	}
	switch (StatusCodeOperator)(*scc.Op) {
	case StatusCodeOperatorNone:
		return func(_ int) bool {
			return false
		}, nil
	case StatusCodeOperatorEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(statusCode int) bool {
			return statusCode == statusCodeInt
		}, nil
	case StatusCodeOperatorNotEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(statusCode int) bool {
			return statusCode != statusCodeInt
		}, nil
	case StatusCodeOperatorLessThan:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(statusCode int) bool {
			return statusCode < statusCodeInt
		}, nil
	case StatusCodeOperatorLessEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(statusCode int) bool {
			return statusCode <= statusCodeInt
		}, nil
	case StatusCodeOperatorGreaterThan:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(statusCode int) bool {
			return statusCode > statusCodeInt
		}, nil
	case StatusCodeOperatorGreaterEqual:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeInt, ok := (*scc.Value).(int)
		if !ok {
			return nil, fmt.Errorf("value must be int")
		}
		return func(statusCode int) bool {
			return statusCode >= statusCodeInt
		}, nil
	case StatusCodeOperatorIn:
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
		return func(statusCode int) bool {
			for _, v := range codes {
				if statusCode == v {
					return true
				}
			}
			return false
		}, nil
	case StatusCodeOperatorNotIn:
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
		return func(statusCode int) bool {
			for _, v := range codes {
				if statusCode == v {
					return false
				}
			}
			return true
		}, nil
	case StatusCodeOperatorBetween:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeVals, ok := (*scc.Value).(statusCodeBetweenVal)
		if !ok {
			return nil, fmt.Errorf("value must be statusCodeBetweenVal")
		}
		return func(statusCode int) bool {
			return statusCode >= statusCodeVals.Min && statusCode <= statusCodeVals.Max
		}, nil
	case StatusCodeOperatorNotBetween:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		statusCodeVals, ok := (*scc.Value).(statusCodeBetweenVal)
		if !ok {
			return nil, fmt.Errorf("value must be statusCodeBetweenVal")
		}
		return func(statusCode int) bool {
			return statusCode < statusCodeVals.Min || statusCode > statusCodeVals.Max
		}, nil
	case StatusCodeOperatorRegex:
		if scc.Value == nil {
			return nil, fmt.Errorf("value is required")
		}
		strV, ok := (*scc.Value).(string)
		if !ok {
			return nil, fmt.Errorf("value must be string")
		}
		preRequireRegex := regexp.MustCompile(strV)
		return func(statusCode int) bool {
			return preRequireRegex.MatchString(strconv.Itoa(statusCode))
		}, nil
	default:
		log.Error(ctx, "unknown operator",
			logger.Value("operator", scc.Op), logger.Value("on", "statusCodeMatherFactory"))
		return nil, fmt.Errorf("unknown operator: %s", *scc.Op)
	}
}

// StatusCodeConditions represents the status code conditions
type StatusCodeConditions []StatusCodeCondition

// StatusCodeConditionsMatcher represents the status code conditions matcher
type StatusCodeConditionsMatcher func(statusCode int) (string, bool)

// MatcherGenerate generates the status code conditions matcher
func (sccs StatusCodeConditions) MatcherGenerate(
	ctx context.Context,
	log logger.Logger,
) (StatusCodeConditionsMatcher, error) {
	matchers := make([]StatusCodeConditionMatcher, 0, len(sccs))
	for _, scc := range sccs {
		matcher, err := scc.MatcherGenerate(ctx, log)
		if err != nil {
			return nil, fmt.Errorf("failed to generate matcher: %w", err)
		}
		matchers = append(matchers, matcher)
	}
	return func(statusCode int) (string, bool) {
		for i, matcher := range matchers {
			if matcher(statusCode) {
				if sccs[i].ID != nil {
					return *sccs[i].ID, true
				}
				return "", true
			}
		}
		return "", false
	}, nil
}
