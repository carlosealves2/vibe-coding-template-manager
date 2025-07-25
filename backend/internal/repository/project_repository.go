package repository

import (
	"context"
	"template-manager-backend/internal/domain"

	"github.com/phuslu/log"
	"gorm.io/gorm"
)

// projectRepository implementa domain.ProjectRepository
type projectRepository struct {
	db *gorm.DB
}

// NewProjectRepository cria uma nova instância do repositório de projetos
func NewProjectRepository(db *gorm.DB) domain.ProjectRepository {
	return &projectRepository{db: db}
}

// Create cria um novo projeto
func (r *projectRepository) Create(ctx context.Context, project *domain.Project) error {
	if err := r.db.WithContext(ctx).Create(project).Error; err != nil {
		log.Error().Err(err).Str("name", project.Name).Msg("failed to insert project")
		return err
	}
	log.Info().Uint("id", project.ID).Msg("project inserted")
	return nil
}

// GetByID busca um projeto por ID
func (r *projectRepository) GetByID(ctx context.Context, id uint) (*domain.Project, error) {
	var project domain.Project
	err := r.db.WithContext(ctx).Preload("Template").First(&project, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

// GetAll busca todos os projetos
func (r *projectRepository) GetAll(ctx context.Context) ([]*domain.Project, error) {
	var projects []*domain.Project
	err := r.db.WithContext(ctx).Preload("Template").Find(&projects).Error
	return projects, err
}

// Update atualiza um projeto existente
func (r *projectRepository) Update(ctx context.Context, project *domain.Project) error {
	return r.db.WithContext(ctx).Save(project).Error
}

// Delete remove um projeto
func (r *projectRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Project{}, id).Error
}

// GetByName busca um projeto por nome
func (r *projectRepository) GetByName(ctx context.Context, name string) (*domain.Project, error) {
	var project domain.Project
	err := r.db.WithContext(ctx).Preload("Template").Where("name = ?", name).First(&project).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}
