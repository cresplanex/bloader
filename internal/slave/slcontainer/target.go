package slcontainer

import (
	"fmt"
	"sync"

	pb "github.com/cresplanex/bloader/gen/pb/cresplanex/bloader/v1"

	"github.com/cresplanex/bloader/internal/config"
	"github.com/cresplanex/bloader/internal/target"
)

// Target represents the target container for the slave node
type Target struct {
	mu *sync.RWMutex
	target.Container
}

// NewTarget creates a new target container for the slave node
func NewTarget() *Target {
	return &Target{
		mu:        &sync.RWMutex{},
		Container: make(map[string]target.Target),
	}
}

// Exists checks if the target exists
func (t Target) Exists(id string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	_, ok := t.Container[id]
	return ok
}

// Add adds a new target to the container
func (t *Target) Add(id string, target target.Target) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Container[id] = target
}

// Remove removes a target from the container
func (t *Target) Remove(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	delete(t.Container, id)
}

// AddFromProto adds a new target from the proto to the container
func (t Target) AddFromProto(id string, pbT *pb.Target) error {
	switch pbT.Type {
	case pb.TargetType_TARGET_TYPE_HTTP:
		t.Add(id, target.Target{
			Type: config.TargetTypeHTTP,
			URL:  pbT.GetHttp().Url,
		})
		return nil
	case pb.TargetType_TARGET_TYPE_UNSPECIFIED:
		return fmt.Errorf("invalid target type: %v", pbT.Type)
	}

	return fmt.Errorf("invalid target type: %v", pbT.Type)
}

// Find finds a target from the container
func (t Target) Find(id string) (target.Target, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	target, ok := t.Container[id]
	return target, ok
}

// TargetResourceRequest represents a target resource request
type TargetResourceRequest struct {
	TargetID string
}
