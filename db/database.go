package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/shiibs/fitness-app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(*gorm.DB) error) *gorm.DB {
	uri := fmt.Sprintf(`
	host=%s user=%s dbname=%s password=%s sslmode=%s port=5432`,
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %e", err)
	}

	log.Info("Connected to database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate: %e", err)
	}

	return db
}
