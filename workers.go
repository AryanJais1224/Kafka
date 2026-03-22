package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func worker(id int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "workers",
	})

	for {
		m, _ := r.ReadMessage(context.Background())
		fmt.Println("worker", id, string(m.Value))
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go worker(i)
	}
	select {}
}
