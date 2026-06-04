// Why: auth routes stay together so public login/register endpoints are easy to find.
// What to do: add public authentication endpoints here, such as refresh token or forgot password.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
)

func RegisterAuthRoutes(r *gin.RouterGroup, h *handler.Handler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}
