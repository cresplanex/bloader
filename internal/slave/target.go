package slave

import (
	"context"
	"fmt"

	"github.com/cresplanex/bloader/internal/runner"
	"github.com/cresplanex/bloader/internal/slave/slcontainer"
	"github.com/cresplanex/bloader/internal/target"
)

// TargetFactor represents the factory for the slave target
type TargetFactor struct {
	target                        *slcontainer.Target
	connectionID                  string
	receiveChanelRequestContainer *slcontainer.ReceiveChanelRequestContainer
	mapper                        *slcontainer.RequestConnectionMapper
}

// Factorize returns the factorized target
func (s *TargetFactor) Factorize(ctx context.Context, targetID string) (target.Target, error) {
	t, ok := s.target.Find(targetID)
	if ok {
		return t, nil
	}

	term := s.receiveChanelRequestContainer.SendTargetResourceRequests(
		ctx,
		s.connectionID,
		s.mapper,
		slcontainer.TargetResourceRequest{
			TargetID: targetID,
		},
	)
	if term == nil {
		return target.Target{}, fmt.Errorf("failed to send target resource request")
	}
	select {
	case <-ctx.Done():
		return target.Target{}, nil
	case <-term:
	}

	t, ok = s.target.Find(targetID)
	if !ok {
		return target.Target{}, fmt.Errorf("target not found: %s", targetID)
	}

	return t, nil
}

var _ runner.TargetFactor = &TargetFactor{}
