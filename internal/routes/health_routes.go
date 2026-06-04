package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
)

func RegisterHealthRoutes(r *gin.Engine, h *handler.Handler) {
	r.GET("/health", h.Health)
}
