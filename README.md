# Kafka User Logger 🚀

This project demonstrates a **Kafka producer and consumer written in Go** using the [segmentio/kafka-go](https://github.com/segmentio/kafka-go) client.

It simulates a **real-time user signup event flow**, perfect for learning Kafka + Go, and preparing for interviews (e.g. Confluent, backend systems, etc.)

---

## 📦 Tech Stack

- **Go 1.22**
- **Kafka** (via Docker)
- **Kafka-Go** client
- **Docker Compose** for local Kafka + Zookeeper setup

---

## 📌 Features

- `producer.go`: Sends signup events to Kafka topic `user-signups`
- `consumer.go`: Listens to `user-signups` and prints new messages
- Clean, minimal, and interview-ready

---

## ▶️ How to Run

### ✅ 1. Start Kafka and Zookeeper

```bash
docker network create kafka-net

docker run -d --name zookeeper --network kafka-net -p 2181:2181 confluentinc/cp-zookeeper:7.4.0

docker run -d --name kafka --network kafka-net -p 9092:9092 \
  -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  -e KAFKA_BROKER_ID=1 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  confluentinc/cp-kafka:7.4.0


---

## 📬 Sample Output

### ▶️ Producer Terminal:

```bash
C:\Projects\kafka-user-logger> go run producer.go
✅ Message sent!

# Kafka User Logger 🚀

![Go Version](https://img.shields.io/badge/go-1.22-blue)
![Last Commit](https://img.shields.io/github/last-commit/Ria36/kafka-user-logger)
![License](https://img.shields.io/github/license/Ria36/kafka-user-logger)
![Stars](https://img.shields.io/github/stars/Ria36/kafka-user-logger?style=social)

Kafka producer and consumer written in Go using the segmentio/kafka-go client.
