package usecase

import (
	"context"
	"errors"
	"template-manager-backend/internal/domain"
)

// TemplateUseCase implementa a lógica de negócio para templates
type TemplateUseCase struct {
	templateRepo domain.TemplateRepository
}

// NewTemplateUseCase cria uma nova instância do use case de templates
func NewTemplateUseCase(templateRepo domain.TemplateRepository) *TemplateUseCase {
	return &TemplateUseCase{
		templateRepo: templateRepo,
	}
}

// CreateTemplate cria um novo template
func (uc *TemplateUseCase) CreateTemplate(ctx context.Context, req *domain.CreateTemplateRequest) (*domain.Template, error) {
	// Verificar se já existe um template com o mesmo nome
	existing, _ := uc.templateRepo.GetByName(ctx, req.Name)
	if existing != nil {
		return nil, errors.New("template with this name already exists")
	}

	template := &domain.Template{
		Name:        req.Name,
		Description: req.Description,
		GitURL:      req.GitURL,
		Language:    req.Language,
		Tags:        req.Tags,
	}

	if err := uc.templateRepo.Create(ctx, template); err != nil {
		return nil, err
	}

	return template, nil
}

// GetTemplate busca um template por ID
func (uc *TemplateUseCase) GetTemplate(ctx context.Context, id uint) (*domain.Template, error) {
	template, err := uc.templateRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, errors.New("template not found")
	}
	return template, nil
}

// GetAllTemplates busca todos os templates
func (uc *TemplateUseCase) GetAllTemplates(ctx context.Context) ([]*domain.Template, error) {
	return uc.templateRepo.GetAll(ctx)
}

// UpdateTemplate atualiza um template existente
func (uc *TemplateUseCase) UpdateTemplate(ctx context.Context, id uint, req *domain.UpdateTemplateRequest) (*domain.Template, error) {
	template, err := uc.templateRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, errors.New("template not found")
	}

	// Atualizar apenas os campos fornecidos
	if req.Name != "" {
		template.Name = req.Name
	}
	if req.Description != "" {
		template.Description = req.Description
	}
	if req.GitURL != "" {
		template.GitURL = req.GitURL
	}
	if req.Language != "" {
		template.Language = req.Language
	}
	if req.Tags != "" {
		template.Tags = req.Tags
	}

	if err := uc.templateRepo.Update(ctx, template); err != nil {
		return nil, err
	}

	return template, nil
}

// DeleteTemplate remove um template
func (uc *TemplateUseCase) DeleteTemplate(ctx context.Context, id uint) error {
	template, err := uc.templateRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if template == nil {
		return errors.New("template not found")
	}

	return uc.templateRepo.Delete(ctx, id)
}
