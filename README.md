# Bag a simple balance algorithm in golang
Bag is stand for Balance Algorithm in Generic

### Example

generic style
```go



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
```



least connection inspired from https://github.com/hlts2/least-connections