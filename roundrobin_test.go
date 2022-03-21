package bag_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/toeydevelopment/bag"
)

func TestRoundRobin(t *testing.T) {

	t.Run("should return error no arguments when no arguments passed", func(t *testing.T) {
		_, err := bag.NewRoundRobin[int]()
		assert.ErrorIs(t, err, bag.ErrNoArguments)
	})

	t.Run("should return error context canceled when context is canceled", func(t *testing.T) {

		lb, err := bag.NewRoundRobin(1, 2, 3)

		assert.NoError(t, err)

		ctx, cancel := context.WithCancel(context.Background())

		cancel()

		_, err = lb.Next(ctx)

		assert.EqualError(t, context.Canceled, err.Error())
	})

	t.Run("should return 1 when invoked next first time", func(t *testing.T) {
		lb, err := bag.NewRoundRobin(1, 2, 3)

		assert.NoError(t, err)

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 1, result)
	})

	t.Run("should return 2 when invoked next second time", func(t *testing.T) {
		lb, err := bag.NewRoundRobin(1, 2, 3)

		assert.NoError(t, err)

		lb.Next(context.TODO())

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 2, result)
	})

	t.Run("should return 3 when invoked next third time", func(t *testing.T) {
		lb, err := bag.NewRoundRobin(1, 2, 3)

		assert.NoError(t, err)

		lb.Next(context.TODO())
		lb.Next(context.TODO())

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 3, result)
	})

}
