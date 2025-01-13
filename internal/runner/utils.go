package runner

import (
	"context"
	"fmt"
	"time"

	"github.com/cresplanex/bloader/internal/logger"
)

//nolint:unparam
func wait(
	ctx context.Context,
	log logger.Logger,
	conf ValidRunner,
	after SleepValueAfter,
	filename string,
) error {
	if v, wait := conf.RetrieveSleepValue(after); wait {
		log.Debug(ctx, "sleeping after execute",
			logger.Value("duration", v))
		fmt.Println("Sleeping For", v, "at", filename, "...")
		select {
		case <-time.After(v):
		case <-ctx.Done():
			log.Info(ctx, "context done while sleeping",
				logger.Value("on", filename))
			return nil
		}
		fmt.Println("Sleeping Complete", "at", filename)
	}

	return nil
}

func validate(
	ctx context.Context,
	eventCaster EventCaster,
	validateFunc func() error,
) error {
	if err := eventCaster.CastEvent(ctx, RunnerEventValidating); err != nil {
		return fmt.Errorf("failed to cast event: %w", err)
	}
	if err := validateFunc(); err != nil {
		return err
	}
	if err := eventCaster.CastEvent(ctx, RunnerEventValidated); err != nil {
		return fmt.Errorf("failed to cast event: %w", err)
	}

	return nil
}
