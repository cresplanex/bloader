package runner

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/cresplanex/bloader/internal/auth"
	"github.com/cresplanex/bloader/internal/executor/httpexec"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/output"
	"github.com/cresplanex/bloader/internal/utils"
)

// OneExecType represents the type of OneExec
type OneExecType string

const (
	// OneExecTypeHTTP represents the HTTP type
	OneExecTypeHTTP OneExecType = "http"
)

// OneExec represents the OneExec runner
type OneExec struct {
	Type    *string         `yaml:"type"`
	Output  OneExecOutput   `yaml:"output"`
	Auth    OneExecAuth     `yaml:"auth"`
	Request *OneExecRequest `yaml:"request"`
}

// ValidOneExec represents the valid OneExec runner
type ValidOneExec struct {
	Type    OneExecType
	Output  []output.Output
	Auth    auth.SetAuthor
	Request ValidOneExecRequest
}

// Validate validates the OneExec
func (r OneExec) Validate(
	ctx context.Context,
	authFactor AuthenticatorFactor,
	outFactor OutputFactor,
	targetFactor TargetFactor,
) (ValidOneExec, error) {
	var oneExecType OneExecType
	if r.Type == nil {
		return ValidOneExec{}, fmt.Errorf("type is required")
	}
	switch OneExecType(*r.Type) {
	case OneExecTypeHTTP:
		oneExecType = OneExecType(*r.Type)
	default:
		return ValidOneExec{}, fmt.Errorf("invalid type value: %s", *r.Type)
	}
	var validOutput []output.Output
	var validAuth auth.SetAuthor
	validOutput, err := r.Output.Validate(ctx, outFactor)
	if err != nil {
		return ValidOneExec{}, fmt.Errorf("failed to validate output: %w", err)
	}
	validAuth, err = r.Auth.Validate(ctx, authFactor)
	if err != nil {
		return ValidOneExec{}, fmt.Errorf("failed to validate auth: %w", err)
	}
	if r.Request == nil {
		return ValidOneExec{}, fmt.Errorf("request is required")
	}
	validRequest, err := r.Request.Validate(ctx, targetFactor)
	if err != nil {
		return ValidOneExec{}, fmt.Errorf("failed to validate request: %w", err)
	}
	return ValidOneExec{
		Type:    oneExecType,
		Output:  validOutput,
		Auth:    validAuth,
		Request: validRequest,
	}, nil
}

// OneExecOutput represents the output configuration for the OneExec runner
type OneExecOutput struct {
	Enabled bool     `yaml:"enabled"`
	IDs     []string `yaml:"ids"`
}

// Validate validates the OneExecOutput
func (o OneExecOutput) Validate(ctx context.Context, outFactor OutputFactor) ([]output.Output, error) {
	if !o.Enabled {
		return nil, nil
	}
	var outputs []output.Output
	for _, id := range o.IDs {
		output, err := outFactor.Factorize(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to factorize output: %w", err)
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

// OneExecAuth represents the auth configuration for the OneExec runner
type OneExecAuth struct {
	Enabled bool    `yaml:"enabled"`
	AuthID  *string `yaml:"auth_id"`
}

// Validate validates the OneExecAuth
func (a OneExecAuth) Validate(ctx context.Context, authFactor AuthenticatorFactor) (auth.SetAuthor, error) {
	if !a.Enabled {
		return nil, nil
	}
	var authID string
	var isDefault bool
	if a.AuthID == nil {
		isDefault = true
	} else {
		authID = *a.AuthID
	}
	auth, err := authFactor.Factorize(
		ctx,
		authID,
		isDefault,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to factorize auth: %w", err)
	}

	return auth, nil
}

// OneExecRequest represents the request configuration for the OneExec runner
type OneExecRequest struct {
	TargetID      *string                `yaml:"target_id"`
	Endpoint      *string                `yaml:"endpoint"`
	Method        *string                `yaml:"method"`
	QueryParam    map[string]any         `yaml:"query_param"`
	PathVariables map[string]string      `yaml:"path_variables"`
	Headers       map[string]any         `yaml:"headers"`
	BodyType      *string                `yaml:"body_type"`
	Body          any                    `yaml:"body"`
	ResponseType  *string                `yaml:"response_type"`
	Data          []ExecRequestData      `yaml:"data"`
	MemoryData    []ExecRequestData      `yaml:"memory_data"`
	StoreData     []ExecRequestStoreData `yaml:"store_data"`
}

// ValidOneExecRequest represents the valid request configuration for the OneExec runner
type ValidOneExecRequest struct {
	URL           string
	Method        string
	QueryParam    map[string]any
	PathVariables map[string]string
	Headers       map[string]any
	BodyType      HTTPRequestBodyType
	Body          any
	ResponseType  string
	Data          ValidExecRequestDataSlice
	MemoryData    ValidExecRequestDataSlice
	StoreData     []ValidExecRequestStoreData
}

// Validate validates the OneExecRequest
func (r OneExecRequest) Validate(ctx context.Context, targetFactor TargetFactor) (ValidOneExecRequest, error) {
	var valid ValidOneExecRequest
	var err error
	if r.TargetID == nil {
		return ValidOneExecRequest{}, fmt.Errorf("target_id is required")
	}
	if r.Endpoint == nil {
		return ValidOneExecRequest{}, fmt.Errorf("endpoint is required")
	}
	var urlRoot string
	tg, err := targetFactor.Factorize(ctx, *r.TargetID)
	if err != nil {
		return ValidOneExecRequest{}, fmt.Errorf("failed to factorize target: %w", err)
	}
	urlRoot = tg.URL
	valid.URL = fmt.Sprintf("%s%s", urlRoot, *r.Endpoint)
	if r.Method == nil {
		return ValidOneExecRequest{}, fmt.Errorf("method is required")
	}
	valid.Method = *r.Method

	valid.QueryParam = r.QueryParam
	valid.PathVariables = r.PathVariables
	valid.Headers = r.Headers
	valid.Body = r.Body
	if r.BodyType == nil {
		valid.BodyType = DefaultHTTPRequestBodyType
	} else {
		switch HTTPRequestBodyType(*r.BodyType) {
		case HTTPRequestBodyTypeJSON, HTTPRequestBodyTypeForm, HTTPRequestBodyTypeMultipart:
			valid.BodyType = HTTPRequestBodyTypeJSON
		default:
			return ValidOneExecRequest{}, fmt.Errorf("invalid body_type value: %s", *r.BodyType)
		}
	}
	if r.ResponseType == nil {
		return ValidOneExecRequest{}, fmt.Errorf("response_type is required")
	}
	valid.ResponseType = *r.ResponseType
	for _, d := range r.Data {
		validData, err := d.Validate()
		if err != nil {
			return ValidOneExecRequest{}, fmt.Errorf("failed to validate data: %w", err)
		}
		valid.Data = append(valid.Data, validData)
	}
	for _, d := range r.MemoryData {
		validData, err := d.Validate()
		if err != nil {
			return ValidOneExecRequest{}, fmt.Errorf("failed to validate memory data: %w", err)
		}
		valid.MemoryData = append(valid.MemoryData, validData)
	}
	for _, d := range r.StoreData {
		validData, err := d.Validate()
		if err != nil {
			return ValidOneExecRequest{}, fmt.Errorf("failed to validate store data: %w", err)
		}
		valid.StoreData = append(valid.StoreData, validData)
	}
	return valid, nil
}

// Run runs the OneExec runner
func (r ValidOneExec) Run(
	ctx context.Context,
	outputRoot string,
	str *sync.Map,
	log logger.Logger,
	store Store,
) error {
	switch r.Type {
	case OneExecTypeHTTP:
		return r.runHTTP(ctx, outputRoot, str, log, store)
	}
	return nil
}

func (r ValidOneExec) runHTTP(
	ctx context.Context,
	outputRoot string,
	str *sync.Map,
	log logger.Logger,
	store Store,
) error {
	req := HTTPRequest{
		Method:        r.Request.Method,
		URL:           r.Request.URL,
		Headers:       r.Request.Headers,
		QueryParams:   r.Request.QueryParam,
		PathVariables: r.Request.PathVariables,
		BodyType:      r.Request.BodyType,
		Body:          r.Request.Body,
		AttachRequestInfo: func(ctx context.Context, req *http.Request) error {
			r.Auth.SetOnRequest(ctx, req)
			return nil
		},
	}
	exe := httpexec.RequestContent[HTTPRequest]{
		Req:          req,
		ResponseType: httpexec.ResponseType(r.Request.ResponseType),
	}

	writers := make([]output.HTTPDataWrite, 0)
	uniqueName := fmt.Sprintf("%s/%s", outputRoot, utils.GenerateUniqueID())
	for _, o := range r.Output {
		writer, closer, err := o.HTTPDataWriteFactory(
			ctx,
			log,
			true,
			uniqueName,
			append(
				[]string{
					"Success",
					"SendDatetime",
					"ReceivedDatetime",
					"Count",
					"ResponseTime",
					"StatusCode",
				},
				r.Request.Data.ExtractHeader()...,
			),
		)
		if err != nil {
			return fmt.Errorf("failed to create writer: %w", err)
		}
		defer closer()
		writers = append(writers, writer)
	}

	resp, err := exe.RequestExecute(ctx, log)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	var data []string
	for _, d := range r.Request.Data {
		result, err := d.Extractor.Extract(resp.Res)
		if err != nil {
			return fmt.Errorf("failed to extract data: %w", err)
		}
		data = append(data, fmt.Sprint(result))
	}
	for _, w := range writers {
		if err := w(ctx, log, append(resp.ToWriteHTTPData().ToSlice(), data...)); err != nil {
			return fmt.Errorf("failed to write data: %w", err)
		}
	}

	for _, d := range r.Request.MemoryData {
		result, err := d.Extractor.Extract(resp.Res)
		if err != nil {
			return fmt.Errorf("failed to extract memory data: %w", err)
		}
		str.Store(d.Key, result)
	}

	if err := store.StoreWithExtractor(ctx, resp.Res, r.Request.StoreData, nil); err != nil {
		return fmt.Errorf("failed to store data: %w", err)
	}

	return nil
}
