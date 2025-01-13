// Package fakeclock provides a fake clock implementation.
package fakeclock

import (
	"sync"
	"time"

	"github.com/cresplanex/bloader/internal/clock"
)

var _ clock.Clock = (*Clock)(nil)

// Clock is a fake clock implementation.
type Clock struct {
	mu  sync.RWMutex
	now time.Time
}

// New returns a new Clock.
func New(t time.Time) *Clock {
	return &Clock{
		now: t,
	}
}

// Now returns the current time.
func (s *Clock) Now() time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.now
}

// SetTime sets the current time.
func (s *Clock) SetTime(t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.now = t
}
