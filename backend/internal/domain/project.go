package domain

import (
	"time"
)

// Project representa um projeto criado a partir de um template
type Project struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	// GitURL pode ser preenchido após a criação do repositório no GitHub
	GitURL     string    `json:"git_url"`
	TemplateID uint      `json:"template_id" gorm:"not null"`
	Template   Template  `json:"template" gorm:"foreignKey:TemplateID"`
	Status     string    `json:"status" gorm:"default:'creating'"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateProjectRequest representa a requisição para criar um projeto
type CreateProjectRequest struct {
	Name       string `json:"name" validate:"required"`
	TemplateID uint   `json:"template_id" validate:"required"`
}

// ProjectStatus representa os possíveis status de um projeto
const (
	ProjectStatusCreating = "creating"
	ProjectStatusReady    = "ready"
	ProjectStatusError    = "error"
)
