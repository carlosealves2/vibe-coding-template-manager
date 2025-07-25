package domain

import (
	"time"
)

// Template representa um template de repositório
type Template struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null;unique"`
	Description string    `json:"description"`
	GitURL      string    `json:"git_url" gorm:"not null"`
	Language    string    `json:"language"`
	Tags        string    `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTemplateRequest representa a requisição para criar um template
type CreateTemplateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	GitURL      string `json:"git_url" validate:"required,url"`
	Language    string `json:"language"`
	Tags        string `json:"tags"`
}

// UpdateTemplateRequest representa a requisição para atualizar um template
type UpdateTemplateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GitURL      string `json:"git_url" validate:"omitempty,url"`
	Language    string `json:"language"`
	Tags        string `json:"tags"`
}
