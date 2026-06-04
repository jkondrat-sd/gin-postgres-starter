// Why: Redis queue code publishes background jobs without coupling services to Asynq details.
// What to do: add enqueue methods here when services need new async work.
package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type Publisher interface {
	EnqueueWelcomeEmail(ctx context.Context, payload WelcomeEmailPayload) error
}

type RedisQueue struct {
	client *asynq.Client
}

func NewRedisQueue(redisAddr string) *RedisQueue {
	return &RedisQueue{
		client: asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr}),
	}
}

func (q *RedisQueue) EnqueueWelcomeEmail(ctx context.Context, payload WelcomeEmailPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal welcome email payload: %w", err)
	}

	task := asynq.NewTask(TaskSendWelcomeEmail, body)
	if _, err := q.client.EnqueueContext(ctx, task, asynq.MaxRetry(3), asynq.Queue("default")); err != nil {
		return fmt.Errorf("enqueue welcome email task: %w", err)
	}

	return nil
}

func (q *RedisQueue) Close() error {
	return q.client.Close()
}
