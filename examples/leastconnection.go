package main

import (
	"context"

	"github.com/toeydevelopment/bag"
)

func LeastConnection() {
	lb, _ := bag.NewLeastConnection(1, 2, 3)

	lb.Next(context.TODO())
}
