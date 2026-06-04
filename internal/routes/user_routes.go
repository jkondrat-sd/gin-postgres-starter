package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("/me", h.Me)
}
