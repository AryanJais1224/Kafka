package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		Topic:        "test-topic",
		RequiredAcks: kafka.RequireAll,
		Compression:  kafka.Snappy,
		MaxAttempts:  3,
	})

	w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte("config message")},
	)
}
