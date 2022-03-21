package main

import "github.com/toeydevelopment/bag"

func Dynamic() {
	var (
		lb  bag.LoadBalancer[int]
		err error
	)

	lb, err = bag.NewLeastConnection(1, 2, 3)

	_ = lb
	_ = err

	lb, err = bag.NewRoundRobin(1, 2, 3)

	_ = lb
	_ = err

}
