package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/config"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
	"github.com/nattakornwarisnarathorn/example-gin/internal/middleware"
)

func SetupRouter(h *handler.Handler, cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	RegisterHealthRoutes(r, h)

	api := r.Group("/api/v1")
	RegisterAuthRoutes(api, h)

	protected := api.Group("/")
	protected.Use(middleware.JWTAuth(cfg.JWTSecret))

	RegisterUserRoutes(protected, h)
	RegisterProjectRoutes(protected, h)

	return r
}
