package github

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"template-manager-backend/internal/domain"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

// gitService implementa domain.GitService
type gitService struct {
	client    *github.Client
	token     string
	username  string
}

// NewGitService cria uma nova instância do serviço Git
func NewGitService(token, username string) domain.GitService {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &gitService{
		client:   client,
		token:    token,
		username: username,
	}
}

// CloneRepository clona um repositório Git
func (s *gitService) CloneRepository(ctx context.Context, gitURL, destPath string) error {
	// Usar git command line para clonar (mais confiável)
	cmd := exec.CommandContext(ctx, "git", "clone", gitURL, destPath)
	return cmd.Run()
}

// CreateRepository cria um novo repositório no GitHub
func (s *gitService) CreateRepository(ctx context.Context, name, description string) (string, error) {
	repo := &github.Repository{
		Name:        github.String(name),
		Description: github.String(description),
		Private:     github.Bool(false),
	}

       owner := ""
       if s.username != "" {
               owner = s.username
       }
       createdRepo, _, err := s.client.Repositories.Create(ctx, owner, repo)
       if err != nil {
               return "", fmt.Errorf("failed to create repository: %w", err)
       }

	return createdRepo.GetCloneURL(), nil
}

// PushToRepository faz push do código local para um repositório remoto
func (s *gitService) PushToRepository(ctx context.Context, localPath, repoURL string) error {
	// Inicializar repositório Git
	cmd := exec.CommandContext(ctx, "git", "init")
	cmd.Dir = localPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to init git: %w", err)
	}

	// Adicionar todos os arquivos
	cmd = exec.CommandContext(ctx, "git", "add", ".")
	cmd.Dir = localPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add files: %w", err)
	}

	// Fazer commit inicial
	cmd = exec.CommandContext(ctx, "git", "commit", "-m", "Initial commit from template")
	cmd.Dir = localPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	// Adicionar remote origin
	cmd = exec.CommandContext(ctx, "git", "remote", "add", "origin", repoURL)
	cmd.Dir = localPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add remote: %w", err)
	}

	// Push para o repositório
	cmd = exec.CommandContext(ctx, "git", "push", "-u", "origin", "main")
	cmd.Dir = localPath
	cmd.Env = append(os.Environ(), fmt.Sprintf("GIT_ASKPASS=echo"), fmt.Sprintf("GIT_USERNAME=%s", s.username), fmt.Sprintf("GIT_PASSWORD=%s", s.token))
	
	return cmd.Run()
}

// ClearGitHistory remove o histórico de commits de um repositório
func (s *gitService) ClearGitHistory(ctx context.Context, repoPath string) error {
	// Remover diretório .git
	gitDir := filepath.Join(repoPath, ".git")
	if err := os.RemoveAll(gitDir); err != nil {
		return fmt.Errorf("failed to remove .git directory: %w", err)
	}

	return nil
}
