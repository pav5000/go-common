package ctxsleep

import (
	"context"
	"time"
)

// Sleep waits specified duration and exits earlier if context was cancelled
// returns nil if it waited the whole specified duration
// returns error if the context was cancelled earlier.
func Sleep(ctx context.Context, d time.Duration) error {
	if d == 0 {
		return nil
	}
	select {
	case <-time.NewTimer(d).C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
