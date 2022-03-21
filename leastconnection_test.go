package bag_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/toeydevelopment/bag"
)

func TestLeastConnection(t *testing.T) {

	t.Run("should return error no arguments when no arguments passed", func(t *testing.T) {
		_, err := bag.NewLeastConnection[int]()
		assert.ErrorIs(t, err, bag.ErrNoArguments)
	})

	t.Run("should return 1 when invoked next first time", func(t *testing.T) {
		lb, err := bag.NewLeastConnection(1, 2, 3)

		assert.NoError(t, err)

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 1, result)
	})

	t.Run("should return 2 when invoked next second time", func(t *testing.T) {
		lb, err := bag.NewLeastConnection(1, 2, 3)

		assert.NoError(t, err)

		lb.Next(context.TODO())

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 2, result)
	})

	t.Run("should return 3 when invoked next third time", func(t *testing.T) {
		lb, err := bag.NewLeastConnection(1, 2, 3)

		assert.NoError(t, err)

		lb.Next(context.TODO())
		lb.Next(context.TODO())

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 3, result)
	})

	t.Run("should return 3 when 3 is leastest usage", func(t *testing.T) {
		lb, err := bag.NewLeastConnection(1, 2, 3)

		assert.NoError(t, err)

		c1, cancel1 := context.WithCancel(context.Background())
		c2, cancel2 := context.WithCancel(context.Background())
		c3, cancel3 := context.WithCancel(context.Background())
		c4, cancel4 := context.WithCancel(context.Background())
		c5, cancel5 := context.WithCancel(context.Background())
		c6, cancel6 := context.WithCancel(context.Background())

		r1, _ := lb.Next(c1) // 1
		r2, _ := lb.Next(c2) // 2
		r3, _ := lb.Next(c3) // 3
		r4, _ := lb.Next(c4) // 1
		r5, _ := lb.Next(c5) // 2
		r6, _ := lb.Next(c6) // 3
		cancel3()

		assert.Equal(t, 1, r1)
		assert.Equal(t, 2, r2)
		assert.Equal(t, 3, r3)
		assert.Equal(t, 1, r4)
		assert.Equal(t, 2, r5)
		assert.Equal(t, 3, r6)

		// sleep for ensure that context sent signal done
		time.Sleep(time.Millisecond)

		result, err := lb.Next(context.TODO())

		assert.NoError(t, err)

		assert.Equal(t, 3, result)

		cancel1()
		cancel2()
		cancel4()
		cancel5()
		cancel6()
	})

}
