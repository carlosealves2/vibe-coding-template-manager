package repository

import (
	"context"
	"template-manager-backend/internal/domain"

	"gorm.io/gorm"
)

// templateRepository implementa domain.TemplateRepository
type templateRepository struct {
	db *gorm.DB
}

// NewTemplateRepository cria uma nova instância do repositório de templates
func NewTemplateRepository(db *gorm.DB) domain.TemplateRepository {
	return &templateRepository{db: db}
}

// Create cria um novo template
func (r *templateRepository) Create(ctx context.Context, template *domain.Template) error {
	return r.db.WithContext(ctx).Create(template).Error
}

// GetByID busca um template por ID
func (r *templateRepository) GetByID(ctx context.Context, id uint) (*domain.Template, error) {
	var template domain.Template
	err := r.db.WithContext(ctx).First(&template, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &template, nil
}

// GetAll busca todos os templates
func (r *templateRepository) GetAll(ctx context.Context) ([]*domain.Template, error) {
	var templates []*domain.Template
	err := r.db.WithContext(ctx).Find(&templates).Error
	return templates, err
}

// Update atualiza um template existente
func (r *templateRepository) Update(ctx context.Context, template *domain.Template) error {
	return r.db.WithContext(ctx).Save(template).Error
}

// Delete remove um template
func (r *templateRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Template{}, id).Error
}

// GetByName busca um template por nome
func (r *templateRepository) GetByName(ctx context.Context, name string) (*domain.Template, error) {
	var template domain.Template
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&template).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &template, nil
}
