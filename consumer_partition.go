package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "test-topic",
		Partition: 0,
	})

	for {
		m, _ := r.ReadMessage(context.Background())
		fmt.Println("P0:", string(m.Value))
	}
}
