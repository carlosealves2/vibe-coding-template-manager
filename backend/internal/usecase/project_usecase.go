package usecase

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"template-manager-backend/internal/domain"

	"github.com/phuslu/log"
)

// ProjectUseCase implementa a lógica de negócio para projetos
type ProjectUseCase struct {
	projectRepo  domain.ProjectRepository
	templateRepo domain.TemplateRepository
	gitService   domain.GitService
}

// NewProjectUseCase cria uma nova instância do use case de projetos
func NewProjectUseCase(
	projectRepo domain.ProjectRepository,
	templateRepo domain.TemplateRepository,
	gitService domain.GitService,
) *ProjectUseCase {
	return &ProjectUseCase{
		projectRepo:  projectRepo,
		templateRepo: templateRepo,
		gitService:   gitService,
	}
}

// CreateProject cria um novo projeto a partir de um template
func (uc *ProjectUseCase) CreateProject(ctx context.Context, req *domain.CreateProjectRequest) (*domain.Project, error) {
	log.Info().Str("name", req.Name).Uint("template_id", req.TemplateID).Msg("creating project")
	// Verificar se já existe um projeto com o mesmo nome
	existing, _ := uc.projectRepo.GetByName(ctx, req.Name)
	if existing != nil {
		log.Warn().Str("name", req.Name).Msg("project already exists")
		return nil, errors.New("project with this name already exists")
	}

	// Verificar se o template existe
	template, err := uc.templateRepo.GetByID(ctx, req.TemplateID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get template")
		return nil, err
	}
	if template == nil {
		log.Warn().Uint("template_id", req.TemplateID).Msg("template not found")
		return nil, errors.New("template not found")
	}

	// Criar o projeto com status "creating"
	project := &domain.Project{
		Name:       req.Name,
		TemplateID: req.TemplateID,
		Status:     domain.ProjectStatusCreating,
	}

	if err := uc.projectRepo.Create(ctx, project); err != nil {
		log.Error().Err(err).Msg("failed to create project in database")
		return nil, err
	}

	log.Info().Uint("project_id", project.ID).Msg("project record created")

	// Processar a criação do projeto em background
	go uc.processProjectCreation(context.Background(), project, template)

	return project, nil
}

// processProjectCreation processa a criação do projeto em background
func (uc *ProjectUseCase) processProjectCreation(ctx context.Context, project *domain.Project, template *domain.Template) {
	log.Info().Uint("project_id", project.ID).Msg("starting project creation")
	tempDir := filepath.Join(os.TempDir(), fmt.Sprintf("template-%d", project.ID))
	defer os.RemoveAll(tempDir)

	// 1. Clonar o repositório template
	if err := uc.gitService.CloneRepository(ctx, template.GitURL, tempDir); err != nil {
		log.Error().Err(err).Msg("failed to clone repository")
		uc.updateProjectStatus(ctx, project.ID, domain.ProjectStatusError)
		return
	}
	log.Info().Msg("repository cloned")

	// 2. Limpar histórico de commits
	if err := uc.gitService.ClearGitHistory(ctx, tempDir); err != nil {
		log.Error().Err(err).Msg("failed to clear git history")
		uc.updateProjectStatus(ctx, project.ID, domain.ProjectStatusError)
		return
	}
	log.Info().Msg("git history cleared")

	// 3. Criar novo repositório no GitHub
	repoURL, err := uc.gitService.CreateRepository(ctx, project.Name, fmt.Sprintf("Project created from template: %s", template.Name))
	if err != nil {
		log.Error().Err(err).Msg("failed to create repository on github")
		uc.updateProjectStatus(ctx, project.ID, domain.ProjectStatusError)
		return
	}
	log.Info().Str("repo_url", repoURL).Msg("repository created")

	// 4. Fazer push para o novo repositório
	if err := uc.gitService.PushToRepository(ctx, tempDir, repoURL); err != nil {
		log.Error().Err(err).Msg("failed to push project")
		uc.updateProjectStatus(ctx, project.ID, domain.ProjectStatusError)
		return
	}
	log.Info().Msg("code pushed to repository")

	// 5. Atualizar o projeto com a URL do repositório e status "ready"
	project.GitURL = repoURL
	project.Status = domain.ProjectStatusReady
	uc.projectRepo.Update(ctx, project)
	log.Info().Uint("project_id", project.ID).Msg("project ready")
}

// updateProjectStatus atualiza apenas o status do projeto
func (uc *ProjectUseCase) updateProjectStatus(ctx context.Context, projectID uint, status string) {
	project, err := uc.projectRepo.GetByID(ctx, projectID)
	if err != nil {
		log.Error().Err(err).Uint("project_id", projectID).Msg("failed to load project for status update")
		return
	}
	project.Status = status
	uc.projectRepo.Update(ctx, project)
	log.Info().Uint("project_id", projectID).Str("status", status).Msg("status updated")
}

// GetProject busca um projeto por ID
func (uc *ProjectUseCase) GetProject(ctx context.Context, id uint) (*domain.Project, error) {
	project, err := uc.projectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, errors.New("project not found")
	}
	return project, nil
}

// GetAllProjects busca todos os projetos
func (uc *ProjectUseCase) GetAllProjects(ctx context.Context) ([]*domain.Project, error) {
	return uc.projectRepo.GetAll(ctx)
}

// DeleteProject remove um projeto
func (uc *ProjectUseCase) DeleteProject(ctx context.Context, id uint) error {
	project, err := uc.projectRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("project not found")
	}

	return uc.projectRepo.Delete(ctx, id)
}
