package domain

import "context"

// TemplateRepository define as operações de persistência para templates
type TemplateRepository interface {
	Create(ctx context.Context, template *Template) error
	GetByID(ctx context.Context, id uint) (*Template, error)
	GetAll(ctx context.Context) ([]*Template, error)
	Update(ctx context.Context, template *Template) error
	Delete(ctx context.Context, id uint) error
	GetByName(ctx context.Context, name string) (*Template, error)
}

// ProjectRepository define as operações de persistência para projetos
type ProjectRepository interface {
	Create(ctx context.Context, project *Project) error
	GetByID(ctx context.Context, id uint) (*Project, error)
	GetAll(ctx context.Context) ([]*Project, error)
	Update(ctx context.Context, project *Project) error
	Delete(ctx context.Context, id uint) error
	GetByName(ctx context.Context, name string) (*Project, error)
}

// GitService define as operações com repositórios Git
type GitService interface {
	CloneRepository(ctx context.Context, gitURL, destPath string) error
	CreateRepository(ctx context.Context, name, description string) (string, error)
	PushToRepository(ctx context.Context, localPath, repoURL string) error
	ClearGitHistory(ctx context.Context, repoPath string) error
}
