package retrier

import (
	"context"
	"errors"
	"time"

	"github.com/pav5000/go-common/ctxsleep"
)

// Retry the function several times unless cb returns nil or retry count reaches specified retryCount
func Retry(ctx context.Context, retryCount int, pause time.Duration, cb func(context.Context) error) error {
	var lastError error
	for i := 0; i <= retryCount; i++ {
		lastError = cb(ctx)
		if lastError == nil {
			return nil
		}
		err := ctxsleep.Sleep(ctx, pause)
		if err != nil {
			return err
		}
	}
	if lastError == nil {
		return nil
	}
	return errors.New("retry attempts exceeded: " + lastError.Error())
}
