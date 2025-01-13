package matcher

import (
	"fmt"
	"strings"
)

// TerminateType represents the type of terminate
type TerminateType string

const (
	// TerminateTypeByContext represents the context type
	TerminateTypeByContext TerminateType = "context"
	// TerminateTypeByCount represents the count type
	TerminateTypeByCount TerminateType = "count"
	// TerminateTypeBySystemError represents the system error type
	TerminateTypeBySystemError TerminateType = "sysError"
	// TerminateTypeByCreateRequestError represents the create request error type
	TerminateTypeByCreateRequestError TerminateType = "createRequestError"
	// TerminateTypeByParseResponseError represents the parse response error type
	TerminateTypeByParseResponseError TerminateType = "parseError"
	// TerminateTypeByWriteError represents the write error type
	TerminateTypeByWriteError TerminateType = "writeError"
	// TerminateTypeByResponseBodyWriteFilterError represents the response body filter error type
	TerminateTypeByResponseBodyWriteFilterError TerminateType = "responseBodyWriteFilterError"
	// TerminateTypeByResponseBodyDataExtractorError represents the response body extractor error type
	TerminateTypeByResponseBodyDataExtractorError TerminateType = "responseBodyDataExtractorError"
	// TerminateTypeByResponseBodyBreakFilterError represents the response body break filter error type
	TerminateTypeByResponseBodyBreakFilterError TerminateType = "responseBodyBreakFilterError"
	// TerminateTypeByTimeout represents the timeout type
	TerminateTypeByTimeout TerminateType = "time"
	// TerminateTypeByResponseBody represents the response body type
	TerminateTypeByResponseBody TerminateType = "responseBody"
	// TerminateTypeByStatusCode represents the status code type
	TerminateTypeByStatusCode TerminateType = "statusCode"
)

// String returns the string representation of the terminate type
func (t TerminateType) String() string {
	return string(t)
}

// TerminateTypeAndParams represents the terminate type and params
type TerminateTypeAndParams struct {
	Type   TerminateType
	Params []string
}

// TerminateTypeAndParamsSlice represents the slice of TerminateTypeAndParams
type TerminateTypeAndParamsSlice []TerminateTypeAndParams

// NewTerminateTypeAndParams creates a new TerminateTypeAndParams
func NewTerminateTypeAndParams(t TerminateType, p []string) TerminateTypeAndParams {
	return TerminateTypeAndParams{
		Type:   t,
		Params: p,
	}
}

// NewTerminateTypeFromString creates a new TerminateTypeAndParams from a string
func NewTerminateTypeFromString(s string) (TerminateTypeAndParams, error) {
	strs := strings.Split(s, "/")
	var params []string
	if len(strs) != 2 {
		params = nil
	} else {
		params = strings.Split(strs[1], ",")
	}
	switch TerminateType(strs[0]) {
	case TerminateTypeByContext:
		return NewTerminateTypeAndParams(TerminateTypeByContext, nil), nil
	case TerminateTypeByCount:
		return NewTerminateTypeAndParams(TerminateTypeByCount, params), nil
	case TerminateTypeBySystemError:
		return NewTerminateTypeAndParams(TerminateTypeBySystemError, nil), nil
	case TerminateTypeByCreateRequestError:
		return NewTerminateTypeAndParams(TerminateTypeByCreateRequestError, nil), nil
	case TerminateTypeByParseResponseError:
		return NewTerminateTypeAndParams(TerminateTypeByParseResponseError, nil), nil
	case TerminateTypeByWriteError:
		return NewTerminateTypeAndParams(TerminateTypeByWriteError, nil), nil
	case TerminateTypeByTimeout:
		return NewTerminateTypeAndParams(TerminateTypeByTimeout, nil), nil
	case TerminateTypeByResponseBody:
		return NewTerminateTypeAndParams(TerminateTypeByResponseBody, params), nil
	case TerminateTypeByStatusCode:
		return NewTerminateTypeAndParams(TerminateTypeByStatusCode, params), nil
	case TerminateTypeByResponseBodyWriteFilterError:
		return NewTerminateTypeAndParams(TerminateTypeByResponseBodyWriteFilterError, params), nil
	case TerminateTypeByResponseBodyDataExtractorError:
		return NewTerminateTypeAndParams(TerminateTypeByResponseBodyDataExtractorError, params), nil
	case TerminateTypeByResponseBodyBreakFilterError:
		return NewTerminateTypeAndParams(TerminateTypeByResponseBodyBreakFilterError, params), nil
	default:
		return TerminateTypeAndParams{}, fmt.Errorf("invalid terminate type: %s", s)
	}
}

// NewTerminateTypeAndParamsSliceFromStringSlice creates a new TerminateTypeAndParams from a string
func NewTerminateTypeAndParamsSliceFromStringSlice(s []string) (TerminateTypeAndParamsSlice, error) {
	var t []TerminateTypeAndParams
	for _, v := range s {
		tt, err := NewTerminateTypeFromString(v)
		if err != nil {
			return nil, err
		}
		t = append(t, tt)
	}
	return t, nil
}

// Match returns true if the terminate type matches the terminate type and params
func (t TerminateTypeAndParams) Match(terminateType TerminateType, terminateParams string) bool {
	if t.Type != terminateType {
		return false
	}
	if t.Params == nil {
		return true
	}
	for _, p := range t.Params {
		if p == terminateParams {
			return true
		}
	}
	return false
}

// Match returns true if the terminate type and params slice matches the terminate type and params
func (t TerminateTypeAndParamsSlice) Match(terminateType TerminateType, terminateParam string) bool {
	for _, tt := range t {
		if tt.Match(terminateType, terminateParam) {
			return true
		}
	}
	return false
}
