package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	w.WriteMessages(context.Background(),
		kafka.Message{Key: []byte("k1"), Value: []byte("hello")},
		kafka.Message{Key: []byte("k2"), Value: []byte("world")},
	)
}
