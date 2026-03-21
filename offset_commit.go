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
		m, _ := r.FetchMessage(context.Background())
		fmt.Println(string(m.Value))
		r.CommitMessages(context.Background(), m)
	}
}
