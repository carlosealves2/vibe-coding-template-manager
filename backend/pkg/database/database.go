package database

import (
	"template-manager-backend/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDatabase cria uma nova conexão com o banco de dados
func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("template_manager.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate das tabelas
	if err := db.AutoMigrate(&domain.Template{}, &domain.Project{}); err != nil {
		return nil, err
	}

	return db, nil
}
