package dto

import (
	"time"

	"github.com/nattakornwarisnarathorn/example-gin/internal/model"
)

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Key         string `json:"key" binding:"required,alphanum"`
	Description string `json:"description"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"omitempty,oneof=active paused archived"`
}

type ProjectResponse struct {
	ID          uint                `json:"id"`
	OwnerID     uint                `json:"owner_id"`
	Name        string              `json:"name"`
	Key         string              `json:"key"`
	Description string              `json:"description"`
	Status      model.ProjectStatus `json:"status"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func ToProjectResponse(project model.Project) ProjectResponse {
	return ProjectResponse{
		ID:          project.ID,
		OwnerID:     project.OwnerID,
		Name:        project.Name,
		Key:         project.Key,
		Description: project.Description,
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func ToProjectResponses(projects []model.Project) []ProjectResponse {
	responses := make([]ProjectResponse, 0, len(projects))
	for _, project := range projects {
		responses = append(responses, ToProjectResponse(project))
	}
	return responses
}
