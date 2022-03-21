package bag

import (
	"context"
	"sync/atomic"
)

type roundrobin[T any] struct {
	things []T
	next   uint32
}

// NewRoundRobin returns RoundRobin implementation(roundrobin).
func NewRoundRobin[T any](things ...T) (*roundrobin[T], error) {

	if len(things) == 0 {
		return nil, ErrNoArguments
	}

	return &roundrobin[T]{
		things: things,
		next:   0,
	}, nil
}

// Next returns next  things
func (r *roundrobin[T]) Next(ctx context.Context) (T, error) {

	select {
	case <-ctx.Done():
		return getZero[T](), ctx.Err()
	default:
	}

	n := atomic.AddUint32(&r.next, 1)
	return r.things[(int(n)-1)%len(r.things)], nil
}
