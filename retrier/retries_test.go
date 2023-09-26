package retrier

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Retry_ZeroRetryCount_ShouldBeOnlyOneTry(t *testing.T) {
	t.Parallel()
	tries := 0

	err := Retry(context.Background(), 0, time.Millisecond, func(ctx context.Context) error {
		tries++
		return errors.New("test")
	})

	assert.Error(t, err)
	assert.Equal(t, 1, tries)
}

func Test_Retry_OneRetryCount_ShouldBeOneMainTry_AndOneRetryAttempt(t *testing.T) {
	t.Parallel()
	tries := 0

	err := Retry(context.Background(), 1, time.Millisecond, func(ctx context.Context) error {
		tries++
		return errors.New("test")
	})

	assert.Error(t, err)
	assert.Equal(t, 2, tries)
}

func Test_Retry_OnContextCancel_ShouldExitFromSleepImmediately(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	start := time.Now()

	err := Retry(ctx, 1, time.Hour, func(ctx context.Context) error {
		return errors.New("test")
	})

	assert.Error(t, err)
	assert.Greater(t, time.Second, time.Since(start))
}

func Test_Retry_StopsRetriesWhenReturnsNil(t *testing.T) {
	t.Parallel()
	tries := 0

	err := Retry(context.Background(), 100, time.Millisecond, func(ctx context.Context) error {
		tries++
		if tries == 5 {
			return nil
		}
		return errors.New("test")
	})

	assert.NoError(t, err)
	assert.Equal(t, 5, tries)
}

func Test_Retry_WhenRetryLimitExceeded_ErrorTextContainsTheLastErrorAndStopReason(t *testing.T) {
	t.Parallel()

	err := Retry(context.Background(), 1, time.Millisecond, func(ctx context.Context) error {
		return errors.New("testerr")
	})

	assert.ErrorContains(t, err, "testerr")
	assert.ErrorContains(t, err, "retry attempts exceeded:")
}
