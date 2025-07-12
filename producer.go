package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "user-signups",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	message := kafka.Message{
		Key:   []byte("user1"),
		Value: []byte("Hello, user signed up at " + time.Now().Format(time.RFC3339)),
	}

	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		fmt.Println("❌ Error sending message:", err)
	} else {
		fmt.Println("✅ Message sent!")
	}
}

