package lib

import (
	"fmt"

	"github.com/drpepperlover0/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=db port=5432 user=postgres password=admin dbname=pgsql sslmode=disable"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create DB error: %w", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("migrate DB error: %w", err)
	}

	return db, nil
}
