package httpexec

import (
	"context"
	"net/http"

	"github.com/cresplanex/bloader/internal/logger"
)

// ExecReq represents the request executor
type ExecReq interface {
	// CreateRequest creates the http.Request object for the query
	CreateRequest(ctx context.Context, log logger.Logger, count int) (*http.Request, error)
}
