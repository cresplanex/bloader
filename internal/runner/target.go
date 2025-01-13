package runner

import (
	"context"
	"fmt"

	"github.com/cresplanex/bloader/internal/target"
)

// TargetFactor represents the target factor
type TargetFactor interface {
	// Factorize returns the factorized target
	Factorize(ctx context.Context, targetID string) (target.Target, error)
}

// LocalTargetFactor represents the local target factor
type LocalTargetFactor struct {
	targets target.Container
}

// NewLocalTargetFactor creates a new local target factor
func NewLocalTargetFactor(targets target.Container) *LocalTargetFactor {
	return &LocalTargetFactor{
		targets: targets,
	}
}

// Factorize returns the factorized target
func (l LocalTargetFactor) Factorize(_ context.Context, targetID string) (target.Target, error) {
	if target, ok := l.targets[targetID]; ok {
		return target, nil
	}
	return target.Target{}, fmt.Errorf("target not found: %s", targetID)
}
