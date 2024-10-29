package service

import (
	"github.com/shiibs/fitness-app/models"
	"gorm.io/gorm"
)

type FoodService struct {
	db *gorm.DB
}

func NewFoodService(db *gorm.DB) FoodService {
	return FoodService{db}
}

func (s *FoodService) LogFoodEntry(foodEntry *models.FoodEntry) error {
	dailyLog, err := GetOrCreateDailyLog(s.db, foodEntry.UserID)
	if err != nil {
		return err
	}
	foodEntry.DailyLogID = dailyLog.ID
	return s.db.Create(foodEntry).Error
}
