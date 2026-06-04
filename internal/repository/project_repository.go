// Why: project repositories keep project database access separate from business rules.
// What to do: add project query methods here, and keep ownership filters close to the SQL/GORM call.
package repository

import (
	"github.com/nattakornwarisnarathorn/example-gin/internal/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *model.Project) error
	ListByOwner(ownerID uint) ([]model.Project, error)
	FindByID(ownerID uint, id uint) (*model.Project, error)
	Update(project *model.Project) error
	Delete(project *model.Project) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(project *model.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) ListByOwner(ownerID uint) ([]model.Project, error) {
	var projects []model.Project
	if err := r.db.Where("owner_id = ?", ownerID).Order("created_at desc").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *projectRepository) FindByID(ownerID uint, id uint) (*model.Project, error) {
	var project model.Project
	if err := r.db.Where("owner_id = ? AND id = ?", ownerID, id).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *projectRepository) Update(project *model.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepository) Delete(project *model.Project) error {
	return r.db.Delete(project).Error
}
