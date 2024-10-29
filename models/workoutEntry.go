package models

type WorkoutEntry struct {
	ID             uint    `json:"id" gorm:"primaryKey"`
	UserID         uint    `json:"user_id"`
	DailyLogID     uint    `json:"daily_log_id" gorm:"foreignKey:DailyLogID"`
	CaloriesBurned int32   `json:"calories_burned"`
	WorkoutType    string  `json:"workout_type"`
	Duration       float32 `json:"duration"`
}
