package main

import (
	"time"

	nats "github.com/nats-io/go-nats"
)

func main() {
	natsURL := "nats://localhost:4222"
	nc, _ := nats.Connect(natsURL)
	for i := 1; i < 1000; i++ {
		nc.Publish("foo", []byte("Hello World"))
		time.Sleep(1000 * time.Millisecond)
	}
}
