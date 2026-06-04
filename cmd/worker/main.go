// Why: this entrypoint runs background jobs separately from the HTTP API server.
// What to do: start this process alongside the API when async jobs should be processed.
package main

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"
	"github.com/nattakornwarisnarathorn/example-gin/internal/worker"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	logger.Init(cfg.AppEnv)
	defer logger.Log.Sync()

	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.RedisAddr},
		asynq.Config{
			Concurrency: cfg.WorkerConcurrency,
			Queues: map[string]int{
				"default": 1,
			},
		},
	)

	mux := asynq.NewServeMux()
	worker.RegisterHandlers(mux)

	logger.Log.Info("worker running", zap.String("redis_addr", cfg.RedisAddr))
	if err := server.Run(mux); err != nil {
		logger.Log.Fatal("worker stopped", zap.Error(err))
	}
}
