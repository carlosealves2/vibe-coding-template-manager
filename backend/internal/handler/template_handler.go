package handler

import (
	"strconv"
	"template-manager-backend/internal/domain"
	"template-manager-backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// TemplateHandler gerencia as requisições HTTP para templates
type TemplateHandler struct {
	templateUseCase *usecase.TemplateUseCase
}

// NewTemplateHandler cria uma nova instância do handler de templates
func NewTemplateHandler(templateUseCase *usecase.TemplateUseCase) *TemplateHandler {
	return &TemplateHandler{
		templateUseCase: templateUseCase,
	}
}

// CreateTemplate cria um novo template
func (h *TemplateHandler) CreateTemplate(c *fiber.Ctx) error {
	var req domain.CreateTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	template, err := h.templateUseCase.CreateTemplate(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(template)
}

// GetTemplate busca um template por ID
func (h *TemplateHandler) GetTemplate(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid template ID",
		})
	}

	template, err := h.templateUseCase.GetTemplate(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(template)
}

// GetAllTemplates busca todos os templates
func (h *TemplateHandler) GetAllTemplates(c *fiber.Ctx) error {
	templates, err := h.templateUseCase.GetAllTemplates(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(templates)
}

// UpdateTemplate atualiza um template existente
func (h *TemplateHandler) UpdateTemplate(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid template ID",
		})
	}

	var req domain.UpdateTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	template, err := h.templateUseCase.UpdateTemplate(c.Context(), uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(template)
}

// DeleteTemplate remove um template
func (h *TemplateHandler) DeleteTemplate(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid template ID",
		})
	}

	if err := h.templateUseCase.DeleteTemplate(c.Context(), uint(id)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
