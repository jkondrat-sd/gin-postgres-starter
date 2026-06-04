// Why: health routes give infrastructure a simple unauthenticated endpoint to check service status.
// What to do: keep this lightweight, and add readiness checks only when the app needs them.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
)

func RegisterHealthRoutes(r *gin.Engine, h *handler.Handler) {
	r.GET("/health", h.Health)
}
