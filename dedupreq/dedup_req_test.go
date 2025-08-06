package dedupreq

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_OneRequest_FallsToBackend(t *testing.T) {
	var count atomic.Int32
	d := New(func(ctx context.Context, key int) (int, error) {
		count.Add(1)
		assert.Equal(t, 12, key)

		return 34, nil
	})

	res, err := d.Request(context.Background(), 12)

	require.NoError(t, err)
	assert.Equal(t, 34, res)
	assert.EqualValues(t, 1, count.Load())
}

func Test_TwoParallelRequestsWithTheSameKey_DoOnlyOneRequestToBackend(t *testing.T) {
	t.Parallel()

	requestsStarted := &sync.WaitGroup{}
	requestsStarted.Add(2)
	var count atomic.Int32
	d := New(func(ctx context.Context, key int) (int, error) {
		requestsStarted.Wait()
		time.Sleep(time.Millisecond * 50)
		count.Add(1)
		assert.Equal(t, 12, key)

		return 34, nil
	})

	goroutinesEnded := &sync.WaitGroup{}
	goroutinesEnded.Add(2)
	for range 2 {
		go func() {
			defer goroutinesEnded.Done()
			requestsStarted.Done()
			res, err := d.Request(context.Background(), 12)
			assert.NoError(t, err)
			assert.Equal(t, 34, res)
		}()
	}
	goroutinesEnded.Wait()

	assert.EqualValues(t, 1, count.Load())
}

func Test_TwoParallelRequestsWithDifferentKeys_DoSeparateRequestsToBackend(t *testing.T) {
	t.Parallel()

	requestsStarted := &sync.WaitGroup{}
	requestsStarted.Add(2)
	var keys []int
	var lock sync.Mutex
	d := New(func(ctx context.Context, key int) (int, error) {
		requestsStarted.Wait()
		time.Sleep(time.Millisecond * 50)
		lock.Lock()
		keys = append(keys, key)
		lock.Unlock()

		return key + 10, nil
	})

	goroutinesEnded := &sync.WaitGroup{}
	goroutinesEnded.Add(2)
	for i := range 2 {
		go func() {
			defer goroutinesEnded.Done()
			requestsStarted.Done()
			res, err := d.Request(context.Background(), i)
			assert.NoError(t, err)
			assert.Equal(t, i+10, res)
		}()
	}
	goroutinesEnded.Wait()

	assert.ElementsMatch(t, []int{0, 1}, keys)
}

func Test_TwoSequentialRequests_MakeTwoBackendRequests(t *testing.T) {
	resQueue := make(chan int, 2)
	resQueue <- 1
	resQueue <- 2
	d := New(func(ctx context.Context, key int) (int, error) {
		assert.Equal(t, 42, key)
		return <-resQueue, nil
	})

	res1, err1 := d.Request(context.Background(), 42)
	res2, err2 := d.Request(context.Background(), 42)

	require.NoError(t, err1)
	require.NoError(t, err2)
	assert.EqualValues(t, 1, res1)
	assert.EqualValues(t, 2, res2)
}
