// Package output provides a set of functions to write the data to the output.
package output

import (
	"context"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/logger"
)

// HTTPDataWrite writes the data to the output
type HTTPDataWrite func(ctx context.Context, log logger.Logger, data []string) error

// Output represents a output to be scanned
type Output interface {
	// HTTPDataWriteFactory returns the HTTPDataWrite function
	HTTPDataWriteFactory(
		ctx context.Context,
		log logger.Logger,
		enabled bool,
		uniqueName string,
		header []string,
	) (HTTPDataWrite, Close, error)
}

// Container is a map of outputs
type Container map[string]Output

// NewContainer creates a new OutputContainer
func NewContainer(env string, cfg config.ValidOutputConfig) Container {
	outputs := make(Container)
	for _, output := range cfg {
		var t Output
		var ok bool
		for _, val := range output.Values {
			if val.Env == env {
				switch val.Type {
				case config.OutputTypeLocal:
					t = NewLocalOutput(val)
				}
				ok = true
				break
			}
		}
		if !ok {
			continue
		}
		outputs[output.ID] = t
	}
	return outputs
}
