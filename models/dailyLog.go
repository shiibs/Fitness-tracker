package models

import "time"

type CalorieBudget struct {
	Breakfast    float32 `json:"breakfast"`
	MorningSnack float32 `json:"morning_snack"`
	Lunch        float32 `json:"lunch"`
	EveningSnack float32 `json:"evening_snack"`
	Dinner       float32 `json:"dinner"`
}

type DailyLog struct {
	ID                   uint          `json:"id" gorm:"primaryKey"`
	UserID               uint          `json:"user_id"`
	WaterEntry           int16         `json:"water_entry"` // glasses of water consumed
	Sleep                float32       `json:"sleep"`       // hours of sleep
	TotalCalorieConsumed float32       `json:"total_calorie_consumed"`
	TDEE                 float32       `json:"tdee"`
	CalorieBurned        int32         `json:"calorie_burned"`
	TargetCaloriesBurned int32         `json:"target_calories_burned"`
	Date                 time.Time     `json:"date"`
	CalorieBudgets       CalorieBudget `json:"calorie_budgets"`
	CaloriesPerMeal      CalorieBudget `json:"calories_per_meal"` // Track actual intake
}
