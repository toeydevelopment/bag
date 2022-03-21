package main

import (
	"context"
	"net/url"

	"github.com/toeydevelopment/bag"
)

type server struct {
	url        url.URL
	serverName string
}

func Server() {
	server1, _ := url.Parse("https://localhost:3000")
	server2, _ := url.Parse("https://localhost:3001")
	server3, _ := url.Parse("https://localhost:3002")

	servers := []server{
		{
			url:        *server1,
			serverName: "server1",
		},
		{
			url:        *server2,
			serverName: "server2",
		},
		{
			url:        *server3,
			serverName: "server3",
		},
	}

	var (
		lb bag.LoadBalancer[server]
	)

	// roundrobin
	lb, _ = bag.NewRoundRobin(servers...)

	lb.Next(context.TODO())

	lb, _ = bag.NewLeastConnection(servers...)

	lb.Next(context.TODO())

}
