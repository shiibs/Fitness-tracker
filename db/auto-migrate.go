package db

import (
	"github.com/shiibs/fitness-app/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.WeightEntry{},
		&models.DailyLog{},
		&models.Food{},
		&models.FoodEntry{},
	)
}
