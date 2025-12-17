package kafka

import (
	"efs-workforce/internal/domain"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

// EventPublisher implements the event publisher interface using Kafka
type EventPublisher struct {
	producer    sarama.SyncProducer
	topicPrefix string
}

// NewEventPublisher creates a new Kafka event publisher
func NewEventPublisher(brokers []string, topicPrefix string) (*EventPublisher, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &EventPublisher{
		producer:    producer,
		topicPrefix: topicPrefix,
	}, nil
}

// Publish publishes a domain event
func (p *EventPublisher) Publish(event *domain.Event) error {
	if p == nil || p.producer == nil {
		log.Printf("[EventPublisher] Kafka producer not initialized, skipping event: Type=%s", event.Type)
		return nil
	}

	// Serialize event payload
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		return err
	}

	// Create topic name: workforce.{event_type}
	topic := p.topicPrefix + "." + event.Type

	// Create message
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(payload),
		Headers: []sarama.RecordHeader{
			{Key: []byte("event_type"), Value: []byte(event.Type)},
		},
	}

	// Send message
	partition, offset, err := p.producer.SendMessage(message)
	if err != nil {
		log.Printf("[EventPublisher] Failed to publish event: %v", err)
		return err
	}

	log.Printf("[EventPublisher] Published event: Type=%s, Topic=%s, Partition=%d, Offset=%d", event.Type, topic, partition, offset)
	return nil
}

// Close closes the Kafka producer
func (p *EventPublisher) Close() error {
	if p.producer != nil {
		return p.producer.Close()
	}
	return nil
}
