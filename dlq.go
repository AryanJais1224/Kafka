package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	ID string `json:"id"`
}

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "g2",
	})

	dlq := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "dlq-topic",
	})

	for {
		m, _ := r.ReadMessage(context.Background())

		var e Event
		err := json.Unmarshal(m.Value, &e)
		if err != nil {
			dlq.WriteMessages(context.Background(), kafka.Message{Value: m.Value})
			continue
		}

		fmt.Println("processed:", e.ID)
	}
}
