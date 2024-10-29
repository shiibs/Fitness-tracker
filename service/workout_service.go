package service

import (
	"github.com/shiibs/fitness-app/models"
	"gorm.io/gorm"
)

type WorkoutService struct {
	db *gorm.DB
}

func NewWorkoutService(db *gorm.DB) WorkoutService {
	return WorkoutService{db}
}

func (s *WorkoutService) LogWorkoutEntry(workoutEntry *models.WorkoutEntry) error {
	dailyLog, err := GetOrCreateDailyLog(s.db, workoutEntry.UserID)
	if err != nil {
		return err
	}

	workoutEntry.DailyLogID = dailyLog.ID
	return s.db.Create(workoutEntry).Error
}
