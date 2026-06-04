// Why: this entrypoint shows how another process can consume Kafka domain events from the API.
// What to do: run this process when you want to react to events such as user.registered.
package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/nattakornwarisnarathorn/example-gin/internal/config"
	"github.com/nattakornwarisnarathorn/example-gin/internal/event"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/logger"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	logger.Init(cfg.AppEnv)
	defer logger.Log.Sync()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: event.ParseBrokers(cfg.KafkaBrokers),
		Topic:   cfg.KafkaUserTopic,
		GroupID: cfg.KafkaConsumerGroup,
	})
	defer reader.Close()

	logger.Log.Info(
		"kafka consumer running",
		zap.String("topic", cfg.KafkaUserTopic),
		zap.String("group", cfg.KafkaConsumerGroup),
	)

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				logger.Log.Info("kafka consumer stopped")
				return
			}
			logger.Log.Error("kafka read failed", zap.Error(err))
			continue
		}

		logger.Log.Info(
			"kafka event consumed",
			zap.String("topic", msg.Topic),
			zap.Int("partition", msg.Partition),
			zap.Int64("offset", msg.Offset),
			zap.ByteString("key", msg.Key),
			zap.ByteString("value", msg.Value),
		)
	}
}
