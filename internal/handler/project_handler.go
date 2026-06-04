// Why: project handlers keep HTTP parsing and response formatting for project endpoints in one file.
// What to do: parse route/body data here, then delegate ownership and persistence rules to ProjectService.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nattakornwarisnarathorn/example-gin/internal/dto"
	"github.com/nattakornwarisnarathorn/example-gin/pkg/response"
)

func (h *Handler) ListProjects(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	projects, err := h.projects.List(userID)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.OK(c, projects)
}

func (h *Handler) CreateProject(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	var req dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	project, err := h.projects.Create(userID, req)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.Created(c, project)
}

func (h *Handler) GetProject(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	projectID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	project, err := h.projects.Get(userID, projectID)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.OK(c, project)
}

func (h *Handler) UpdateProject(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	projectID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var req dto.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	project, err := h.projects.Update(userID, projectID, req)
	if err != nil {
		writeServiceError(c, err)
		return
	}

	response.OK(c, project)
}

func (h *Handler) DeleteProject(c *gin.Context) {
	userID, ok := userIDFromContext(c)
	if !ok {
		return
	}

	projectID, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	if err := h.projects.Delete(userID, projectID); err != nil {
		writeServiceError(c, err)
		return
	}

	response.Message(c, http.StatusOK, "project deleted")
}
