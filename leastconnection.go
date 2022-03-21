package bag

import (
	"context"
	"sync"
)

type conn[T any] struct {
	thing T
	cnt   int
}

type leastConnections[T any] struct {
	conns []conn[T]
	mu    *sync.Mutex
}

// NewLeastConnection initializes a new instance of leastConnection
func NewLeastConnection[T any](things ...T) (*leastConnections[T], error) {

	if len(things) == 0 {
		return nil, ErrNoArguments
	}

	conns := make([]conn[T], len(things))
	for i := range conns {
		conns[i] = conn[T]{
			thing: things[i],
			cnt:   0,
		}
	}

	return &leastConnections[T]{
		conns: conns,
		mu:    new(sync.Mutex),
	}, nil
}

func (lc *leastConnections[T]) Next(ctx context.Context) (T, error) {

	select {
	case <-ctx.Done():
		return getZero[T](), ctx.Err()
	default:
	}

	var (
		min = -1
		idx int
	)

	lc.mu.Lock()

	for i, conn := range lc.conns {
		if min == -1 || conn.cnt < min {
			min = conn.cnt
			idx = i
		}
	}
	lc.conns[idx].cnt++

	lc.mu.Unlock()

	go func() {
		// wait for the context to be done
		<-ctx.Done()

		lc.mu.Lock()
		lc.conns[idx].cnt--
		lc.mu.Unlock()
	}()

	return lc.conns[idx].thing, nil
}
