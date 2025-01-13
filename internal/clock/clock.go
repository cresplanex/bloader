// Package clock provides the clock interface for the application.
package clock

import "time"

// Clock the interface for the clock
type Clock interface {
	// Now returns the current time.
	Now() time.Time
}

var clk = &RealClock{}

// New creates a new Clock
func New() Clock {
	return clk
}

// RealClock the real clock
type RealClock struct{}

// Now returns the current time.
func (s *RealClock) Now() time.Time {
	return time.Now()
}
