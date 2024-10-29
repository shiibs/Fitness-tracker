package service

import (
	"time"

	"github.com/shiibs/fitness-app/models"
	"gorm.io/gorm"
)

func GetOrCreateDailyLog(db *gorm.DB, userID uint) (*models.DailyLog, error) {
	var dailyLog models.DailyLog
	today := time.Now().Format("2006-01-02")
	err := db.Where("user_id = ? AND date = ?", userID, today).First(&dailyLog).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		dailyLog = models.DailyLog{UserID: userID, Date: today}
		if err := db.Create(&dailyLog).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}
	return &dailyLog, nil
}
