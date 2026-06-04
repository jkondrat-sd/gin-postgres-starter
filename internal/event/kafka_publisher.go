// Why: Kafka publishers broadcast domain events without coupling services to kafka-go details.
// What to do: add publish methods here when services need to emit new event types.
package event

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type Publisher interface {
	PublishUserRegistered(ctx context.Context, payload UserRegisteredEvent) error
}

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewKafkaPublisher(brokersCSV string, topic string) *KafkaPublisher {
	return &KafkaPublisher{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(ParseBrokers(brokersCSV)...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireOne,
			WriteTimeout: 3 * time.Second,
		},
	}
}

func (p *KafkaPublisher) PublishUserRegistered(ctx context.Context, payload UserRegisteredEvent) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal user registered event: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("user:%d", payload.UserID)),
		Value: body,
		Headers: []kafka.Header{
			{Key: "event_type", Value: []byte(EventUserRegistered)},
		},
		Time: time.Now(),
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("publish user registered event: %w", err)
	}

	return nil
}

func (p *KafkaPublisher) Close() error {
	return p.writer.Close()
}

func ParseBrokers(brokersCSV string) []string {
	parts := strings.Split(brokersCSV, ",")
	brokers := make([]string, 0, len(parts))
	for _, part := range parts {
		broker := strings.TrimSpace(part)
		if broker != "" {
			brokers = append(brokers, broker)
		}
	}

	if len(brokers) == 0 {
		return []string{"localhost:9092"}
	}

	return brokers
}
