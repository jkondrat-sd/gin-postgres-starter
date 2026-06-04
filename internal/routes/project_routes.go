// Why: project routes map project URLs to project handlers without containing business logic.
// What to do: add project-related endpoints here, such as members, tasks, or archived project views.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/handler"
)

func RegisterProjectRoutes(r *gin.RouterGroup, h *handler.Handler) {
	projects := r.Group("/projects")
	{
		projects.GET("", h.ListProjects)
		projects.POST("", h.CreateProject)
		projects.GET("/:id", h.GetProject)
		projects.PUT("/:id", h.UpdateProject)
		projects.DELETE("/:id", h.DeleteProject)
	}
}
