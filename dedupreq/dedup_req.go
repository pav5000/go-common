package dedupreq

import (
	"context"
	"errors"
	"sync"
)

var errNilCallback = errors.New("nil callback")

// Deduper makes sure that only one parallel
// request with the same key can be done at the same time.
// Can be used to protect backend from many duplicated requests
// when cache is empty.
type Deduper[Key comparable, Res any] struct {
	lock             sync.Mutex
	inFlightRequests map[Key]*Request[Key, Res]
	callback         func(context.Context, Key) (Res, error)
}

type Request[Key comparable, Res any] struct {
	res  Res
	err  error
	done chan (struct{})
}

func New[Key comparable, Res any](callback func(context.Context, Key) (Res, error)) *Deduper[Key, Res] {
	return &Deduper[Key, Res]{
		callback:         callback,
		inFlightRequests: make(map[Key]*Request[Key, Res]),
	}
}

// Request if there is no request with this key is in progress, makes request.
// If there is a request with this key in progress, waits it to complete and returns it's result.
// If the master request's context is cancelled, all waiting clients will get an error.
// If the waiting request's context is cancelled, only this waiting client will get an error.
func (d *Deduper[Key, Res]) Request(ctx context.Context, key Key) (Res, error) {
	d.lock.Lock()
	request, ok := d.inFlightRequests[key]
	if !ok {
		request = &Request[Key, Res]{
			done: make(chan struct{}),
		}
		d.inFlightRequests[key] = request
	}
	d.lock.Unlock()

	// waiting for results from the other request with the same key
	if ok {
		select {
		case <-ctx.Done():
			var res Res

			return res, ctx.Err()
		case <-request.done:
			return request.res, request.err
		}
	}

	// making request by ourselves
	if d.callback == nil {
		var res Res

		return res, errNilCallback
	}
	request.res, request.err = d.callback(ctx, key)
	close(request.done)

	return request.res, request.err
}
