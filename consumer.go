package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "user-signups",
		GroupID:   "go-consumer-group-1",
		MinBytes:  1,
		MaxBytes:  10e6,
	})

	defer reader.Close()
	fmt.Println("🔁 Listening for messages on 'user-signups' topic...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Error: %v\n", err)
			continue
		}
		fmt.Printf("📬 Received: key = %s, value = %s\n", msg.Key, msg.Value)
	}
}

