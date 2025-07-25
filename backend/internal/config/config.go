package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config representa a configuração da aplicação
type Config struct {
	Port           string
	GitHubToken    string
	GitHubUsername string
}

// LoadConfig carrega a configuração da aplicação
func LoadConfig() (*Config, error) {
	// Carregar variáveis de ambiente do arquivo .env se existir
	godotenv.Load()

	config := &Config{
		Port:           getEnv("PORT", "8080"),
		GitHubToken:    getEnv("GITHUB_TOKEN", ""),
		GitHubUsername: getEnv("GITHUB_USERNAME", ""),
	}

	return config, nil
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
