package fakeclock_test

import (
	"testing"
	"time"

	"github.com/cresplanex/bloader/internal/clock/fakeclock"
)

// TestClock tests the fake clock implementation.
func TestClock(t *testing.T) {
	ts := time.Date(2023, 1, 2, 12, 34, 56, 0, time.UTC)
	clk := fakeclock.New(ts)

	t.Run("Now", func(tt *testing.T) {
		now := clk.Now()
		if !now.Equal(ts) {
			tt.Errorf("expected %q, got %q", ts, now)
		}
	})

	t.Run("SetTime", func(tt *testing.T) {
		newTS := ts.Add(1 * time.Hour)
		clk.SetTime(newTS)

		now := clk.Now()
		if !now.Equal(newTS) {
			tt.Errorf("expected %q, got %q", newTS, now)
		}
	})
}
