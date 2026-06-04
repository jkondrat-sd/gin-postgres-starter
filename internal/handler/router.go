package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/middleware"
)

func SetupRouter(h *Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	r.GET("/health", h.Health)

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.JWTAuth(h.cfg.JWTSecret))
		{
			protected.GET("/me", h.Me)

			projects := protected.Group("/projects")
			{
				projects.GET("", h.ListProjects)
				projects.POST("", h.CreateProject)
				projects.GET("/:id", h.GetProject)
				projects.PUT("/:id", h.UpdateProject)
				projects.DELETE("/:id", h.DeleteProject)
			}
		}
	}

	return r
}
