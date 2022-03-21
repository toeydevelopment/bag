package bag

import (
	"context"
	"errors"
)

var (
	ErrNoArguments = errors.New("LoadBalancer: no arguments")
)

// LoadBalancer a simple load balancer for things that want to balance usage
type LoadBalancer[T any] interface {
	Next(ctx context.Context) (T, error)
}

// getZero returns zero value of T.
func getZero[T any]() T {
	var result T
	return result
}
