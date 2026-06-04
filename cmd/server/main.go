package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
	"github.com/nattakornwarisnarathorn/example-gin/internal/repository"
	"github.com/nattakornwarisnarathorn/example-gin/internal/routes"
	"github.com/nattakornwarisnarathorn/example-gin/internal/service"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/database"
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

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		logger.Log.Fatal("database connection failed", zap.Error(err))
	}

	if err := database.AutoMigrate(db); err != nil {
		logger.Log.Fatal("database migration failed", zap.Error(err))
	}

	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)

	authService := service.NewAuthService(userRepo, cfg)
	userService := service.NewUserService(userRepo)
	projectService := service.NewProjectService(projectRepo)

	apiHandler := handler.NewHandler(handler.Dependencies{
		Config:         cfg,
		AuthService:    authService,
		UserService:    userService,
		ProjectService: projectService,
	})

	router := routes.SetupRouter(apiHandler, cfg)

	addr := ":" + cfg.AppPort
	logger.Log.Info("server running", zap.String("addr", addr))
	if err := router.Run(addr); err != nil {
		logger.Log.Fatal("server stopped", zap.Error(err))
	}
}
