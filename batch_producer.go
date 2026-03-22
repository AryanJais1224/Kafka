package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "test-topic",
		BatchSize: 10,
	})

	var msgs []kafka.Message
	for i := 0; i < 5; i++ {
		msgs = append(msgs, kafka.Message{
			Key:   []byte(fmt.Sprintf("k%d", i)),
			Value: []byte("data"),
		})
	}

	w.WriteMessages(context.Background(), msgs...)
}
