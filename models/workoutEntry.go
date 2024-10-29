package models

type WorkoutEntry struct {
	ID           uint  `json:"id" gorm:"primaryKey"`
	UserID       uint  `json:"user_id"`
	DailyLogID   uint  `json:"daily_log_id" gorm:"foreignKey:DailyLogID"`
	CalorieCount int32 `json:"calorie_count"`
}
