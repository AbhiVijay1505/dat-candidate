package utils

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// PublishMessage sends a message to the specified Kafka topic
func PublishMessage(topic string, key string, value interface{}) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
		return err
	}
	defer producer.Close()

	// Serialize value to JSON
	valueBytes, err := json.Marshal(value)
	if err != nil {
		log.Printf("Failed to serialize value: %v", err)
		return err
	}

	// Create Kafka message
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          valueBytes,
	}

	// Publish the message
	if err := producer.Produce(message, nil); err != nil {
		log.Printf("Failed to produce message: %v", err)
		return err
	}

	// Wait for message to be delivered
	producer.Flush(15 * 1000)
	log.Printf("Message published to Kafka topic %s: %s", topic, valueBytes)
	return nil
}
