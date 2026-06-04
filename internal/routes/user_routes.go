// Why: user routes group authenticated user endpoints separately from auth endpoints.
// What to do: add profile/account routes here when they belong to the current user.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("/me", h.Me)
}
