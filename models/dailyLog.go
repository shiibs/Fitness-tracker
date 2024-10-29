package models

// type CalorieBudget struct {
// 	Breakfast    float32 `json:"breakfast"`
// 	MorningSnack float32 `json:"morning_snack"`
// 	Lunch        float32 `json:"lunch"`
// 	EveningSnack float32 `json:"evening_snack"`
// 	Dinner       float32 `json:"dinner"`
// }

type DailyLog struct {
	ID                   uint           `json:"id" gorm:"primaryKey"`
	UserID               uint           `json:"user_id"`
	WaterEntries         int32          `json:"water_entries"`
	FoodEntries          []FoodEntry    `json:"food_entries" gorm:"foreignKey:DailyLogID"`
	WorkoutEntries       []WorkoutEntry `json:"workout_entries" gorm:"foreignKey:DailyLogID"`
	Sleep                float32        `json:"sleep"`
	TotalCalorieConsumed float32        `json:"total_calorie_consumed"`
	TDEE                 float32        `json:"tdee"`
	CalorieBurned        int32          `json:"calorie_burned"`
	CaloriesBurnedTarget int32          `json:"calories_burn_target"`
	Date                 string         `json:"date"`
	WaterTarget          int32          `json:"water_target"`
	// CalorieBudgets       []CalorieBudget `json:"calorie_budgets"`
	// CaloriesPerMeal      []CalorieBudget `json:"calories_per_meal"`
}
