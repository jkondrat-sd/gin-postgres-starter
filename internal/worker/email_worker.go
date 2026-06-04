// Why: email workers consume email-related background jobs outside the HTTP request path.
// What to do: replace log-only examples here with real email provider calls when needed.
package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/nattakornwarisnarathorn/example-gin/internal/queue"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/logger"
	"go.uber.org/zap"
)

func RegisterHandlers(mux *asynq.ServeMux) {
	mux.HandleFunc(queue.TaskSendWelcomeEmail, HandleWelcomeEmail)
}

func HandleWelcomeEmail(ctx context.Context, task *asynq.Task) error {
	var payload queue.WelcomeEmailPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("decode welcome email payload: %w", err)
	}

	logger.Log.Info(
		"processed welcome email job",
		zap.Uint("user_id", payload.UserID),
		zap.String("email", payload.Email),
		zap.String("name", payload.Name),
	)

	return nil
}
