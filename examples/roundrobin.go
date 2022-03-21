package main

import (
	"context"

	"github.com/toeydevelopment/bag"
)

func RoundRobin() {
	lb, _ := bag.NewRoundRobin(1, 2, 3)

	lb.Next(context.TODO())
}
