package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "g1",
	})

	for {
		m, _ := r.ReadMessage(context.Background())
		fmt.Println(string(m.Key), string(m.Value), m.Offset)
	}
}
