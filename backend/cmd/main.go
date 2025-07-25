package main

import (
	"github.com/phuslu/log"
	"template-manager-backend/internal/config"
	"template-manager-backend/internal/handler"
	"template-manager-backend/internal/repository"
	"template-manager-backend/internal/usecase"
	"template-manager-backend/pkg/database"
	"template-manager-backend/pkg/github"
	appLogger "template-manager-backend/pkg/logger"

	"github.com/gofiber/fiber/v2"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Configure application logger
	appLogger.Init()

	// Carregar configuração
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// Conectar ao banco de dados
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// Inicializar repositórios
	templateRepo := repository.NewTemplateRepository(db)
	projectRepo := repository.NewProjectRepository(db)

	// Inicializar serviços
	gitService := github.NewGitService(cfg.GitHubToken, cfg.GitHubUsername)

	// Inicializar use cases
	templateUseCase := usecase.NewTemplateUseCase(templateRepo)
	projectUseCase := usecase.NewProjectUseCase(projectRepo, templateRepo, gitService)

	// Inicializar handlers
	templateHandler := handler.NewTemplateHandler(templateUseCase)
	projectHandler := handler.NewProjectHandler(projectUseCase)

	// Criar aplicação Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middlewares
	app.Use(fiberlogger.New())

	// CORS básico
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin,Content-Type,Accept,Authorization")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	// Rotas da API
	api := app.Group("/api/v1")

	// Rotas de templates
	templates := api.Group("/templates")
	templates.Post("/", templateHandler.CreateTemplate)
	templates.Get("/", templateHandler.GetAllTemplates)
	templates.Get("/:id", templateHandler.GetTemplate)
	templates.Put("/:id", templateHandler.UpdateTemplate)
	templates.Delete("/:id", templateHandler.DeleteTemplate)

	// Rotas de projetos
	projects := api.Group("/projects")
	projects.Post("/", projectHandler.CreateProject)
	projects.Get("/", projectHandler.GetAllProjects)
	projects.Get("/:id", projectHandler.GetProject)
	projects.Delete("/:id", projectHandler.DeleteProject)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Iniciar servidor
	log.Info().Str("port", cfg.Port).Msg("Server starting")
	log.Fatal().Err(app.Listen(":" + cfg.Port)).Msg("server exited")
}
