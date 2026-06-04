// Why: project service owns project business rules, including ownership checks and status changes.
// What to do: add project validation and workflow behavior here, keeping database queries in repositories.
package service

import (
	"errors"
	"strings"

	"github.com/nattakornwarisnarathorn/example-gin/internal/dto"
	"github.com/nattakornwarisnarathorn/example-gin/internal/model"
	"github.com/nattakornwarisnarathorn/example-gin/internal/repository"
	"gorm.io/gorm"
)

type ProjectService struct {
	projects repository.ProjectRepository
}

func NewProjectService(projects repository.ProjectRepository) *ProjectService {
	return &ProjectService{projects: projects}
}

func (s *ProjectService) Create(ownerID uint, req dto.CreateProjectRequest) (*dto.ProjectResponse, error) {
	project := model.Project{
		OwnerID:     ownerID,
		Name:        req.Name,
		Key:         strings.ToUpper(req.Key),
		Description: req.Description,
		Status:      model.ProjectStatusActive,
	}

	if err := s.projects.Create(&project); err != nil {
		return nil, err
	}

	res := dto.ToProjectResponse(project)
	return &res, nil
}

func (s *ProjectService) List(ownerID uint) ([]dto.ProjectResponse, error) {
	projects, err := s.projects.ListByOwner(ownerID)
	if err != nil {
		return nil, err
	}
	return dto.ToProjectResponses(projects), nil
}

func (s *ProjectService) Get(ownerID uint, projectID uint) (*dto.ProjectResponse, error) {
	project, err := s.projects.FindByID(ownerID, projectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	res := dto.ToProjectResponse(*project)
	return &res, nil
}

func (s *ProjectService) Update(ownerID uint, projectID uint, req dto.UpdateProjectRequest) (*dto.ProjectResponse, error) {
	project, err := s.projects.FindByID(ownerID, projectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if req.Name != "" {
		project.Name = req.Name
	}
	if req.Description != "" {
		project.Description = req.Description
	}
	if req.Status != "" {
		project.Status = model.ProjectStatus(req.Status)
	}

	if err := s.projects.Update(project); err != nil {
		return nil, err
	}

	res := dto.ToProjectResponse(*project)
	return &res, nil
}

func (s *ProjectService) Delete(ownerID uint, projectID uint) error {
	project, err := s.projects.FindByID(ownerID, projectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}

	return s.projects.Delete(project)
}
