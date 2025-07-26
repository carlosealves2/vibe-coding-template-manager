package handler

import (
	"bufio"
	"fmt"
	"strconv"
	"template-manager-backend/internal/domain"
	"template-manager-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// ProjectHandler gerencia as requisições HTTP para projetos
type ProjectHandler struct {
	projectUseCase *usecase.ProjectUseCase
}

// NewProjectHandler cria uma nova instância do handler de projetos
func NewProjectHandler(projectUseCase *usecase.ProjectUseCase) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}

// CreateProject cria um novo projeto
func (h *ProjectHandler) CreateProject(c *fiber.Ctx) error {
	var req domain.CreateProjectRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	project, err := h.projectUseCase.CreateProject(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(project)
}

// GetProject busca um projeto por ID
func (h *ProjectHandler) GetProject(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID",
		})
	}

	project, err := h.projectUseCase.GetProject(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(project)
}

// GetAllProjects busca todos os projetos
func (h *ProjectHandler) GetAllProjects(c *fiber.Ctx) error {
	projects, err := h.projectUseCase.GetAllProjects(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(projects)
}

// DeleteProject remove um projeto
func (h *ProjectHandler) DeleteProject(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid project ID",
		})
	}

	if err := h.projectUseCase.DeleteProject(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// StreamLogs envia os logs de criação do projeto via SSE.
func (h *ProjectHandler) StreamLogs(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		ch := h.projectUseCase.SubscribeLogs(uint(id))
		for msg := range ch {
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.Flush()
		}
	})
	return nil
}
