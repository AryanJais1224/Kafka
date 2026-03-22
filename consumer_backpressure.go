package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "slow",
	})

	for {
		m, _ := r.ReadMessage(context.Background())
		time.Sleep(2 * time.Second)
		fmt.Println("slow:", string(m.Value))
	}
}
