package httpexec

import (
	"context"
	"strconv"
	"time"

	"github.com/cresplanex/bloader/internal/logger"
)

// ResponseContent represents the response content
type ResponseContent struct {
	Success         bool
	StartTime       time.Time
	EndTime         time.Time
	Res             any
	Count           int
	ByteResponse    []byte
	ResponseTime    int64
	StatusCode      int
	ReqCreateHasErr bool
	ParseResHasErr  bool
	HasSystemErr    bool
	WithCountLimit  bool
}

// ToWriteHTTPData converts the ResponseContent to WriteHTTPData
func (r ResponseContent) ToWriteHTTPData() WriteHTTPData {
	return WriteHTTPData{
		Success:          r.Success,
		SendDatetime:     r.StartTime.Format(time.RFC3339Nano),
		ReceivedDatetime: r.EndTime.Format(time.RFC3339Nano),
		Count:            r.Count,
		ResponseTime:     int(r.ResponseTime),
		StatusCode:       strconv.Itoa(r.StatusCode),
	}
}

// WriteHTTPData represents the data to be written
type WriteHTTPData struct {
	Success          bool
	SendDatetime     string
	ReceivedDatetime string
	Count            int
	ResponseTime     int
	StatusCode       string
}

// ToSlice converts the WriteHTTPData to a slice
func (d WriteHTTPData) ToSlice() []string {
	return []string{
		strconv.FormatBool(d.Success),
		d.SendDatetime,
		d.ReceivedDatetime,
		strconv.Itoa(d.Count),
		strconv.Itoa(d.ResponseTime),
		d.StatusCode,
	}
}

// ResponseType represents the response type
type ResponseType string

const (
	// ResponseTypeJSON represents the JSON response type
	ResponseTypeJSON ResponseType = "json"
	// ResponseTypeXML represents the XML response type
	ResponseTypeXML ResponseType = "xml"
	// ResponseTypeYAML represents the YAML response type
	ResponseTypeYAML ResponseType = "yaml"
	// ResponseTypeText represents the text response type
	ResponseTypeText ResponseType = "text"
	// ResponseTypeHTML represents the HTML response type
	ResponseTypeHTML ResponseType = "html"
)

// RequestExecutor represents the request executor
type RequestExecutor interface {
	// RequestExecute executes the request
	RequestExecute(ctx context.Context, log logger.Logger) (ResponseContent, error)
}

// MassRequestExecutor represents the request executor
type MassRequestExecutor interface {
	// MassRequestExecute executes the request
	MassRequestExecute(ctx context.Context, log logger.Logger) error
}
