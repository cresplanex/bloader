package output

import (
	"context"
	"encoding/csv"
	"fmt"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/utils"
)

// LocalOutput represents the local output service
type LocalOutput struct {
	// Format of the output
	Format config.OutputFormat
	// BasePath of the output
	BasePath string
}

// NewLocalOutput creates a new LocalOutput
func NewLocalOutput(cfg config.ValidOutputRespectiveValueConfig) LocalOutput {
	return LocalOutput{
		Format:   cfg.Format,
		BasePath: cfg.BasePath,
	}
}

// HTTPDataWriteFactory returns the HTTPDataWrite function
func (o LocalOutput) HTTPDataWriteFactory(
	ctx context.Context,
	log logger.Logger,
	enabled bool,
	uniqueName string,
	header []string,
) (HTTPDataWrite, Close, error) {
	var filePath string
	switch o.Format {
	case config.OutputFormatCSV:
		filePath = fmt.Sprintf("%s/%s.csv", o.BasePath, uniqueName)
	default:
		return nil, nil, fmt.Errorf("unsupported output format: %s", o.Format)
	}
	f, err := utils.CreateFileWithDir(filePath)
	if err != nil {
		log.Error(ctx, "failed to create file",
			logger.Value("error", err))
		return nil, nil, fmt.Errorf("failed to create file: %w", err)
	}
	switch o.Format {
	case config.OutputFormatCSV:
		writer := csv.NewWriter(f)
		if err := writer.Write(header); err != nil {
			log.Error(ctx, "failed to write header",
				logger.Value("error", err))
			return nil, nil, fmt.Errorf("failed to write header: %w", err)
		}
		writer.Flush()
	}
	return func(
			ctx context.Context,
			log logger.Logger,
			data []string,
		) error {
			if !enabled {
				return nil
			}
			switch o.Format {
			case config.OutputFormatCSV:
				writer := csv.NewWriter(f)
				log.Debug(ctx, "Writing data to csv",
					logger.Value("data", data))
				if err := writer.Write(data); err != nil {
					log.Error(ctx, "failed to write data to csv",
						logger.Value("error", err))
				}
				writer.Flush()
				return nil
			}
			return fmt.Errorf("unsupported output format: %s", o.Format)
		}, func() error {
			return f.Close()
		}, nil
}

var _ Output = LocalOutput{}
