package services

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"payment-service/models"
)

// KafkaProducerInterface defines methods for Kafka publishing
type KafkaProducerInterface interface {
	Publish(topic string, event models.PaymentEvent) error
}

// KafkaProducer implements KafkaProducerInterface
type KafkaProducer struct {
	producer *kafka.Producer
}

// NewKafkaProducer initializes a Kafka producer
func NewKafkaProducer(brokers string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v", err)
		return nil, err
	}
	return &KafkaProducer{producer: p}, nil
}

// Publish sends a message to Kafka
func (kp *KafkaProducer) Publish(topic string, event models.PaymentEvent) error {
	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return err
	}

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err = kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		log.Printf("Failed to publish message to Kafka: %v", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	log.Printf("Successfully published event to topic %s: %s", topic, string(message))
	return nil
}
