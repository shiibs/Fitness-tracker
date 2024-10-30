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
	ID                   uint           `json:"id" gorm:"primaryKey"`
	UserID               uint           `json:"user_id"`
	WaterEntries         int32          `json:"water_entries"`
	FoodEntries          []FoodEntry    `json:"food_entries" gorm:"foreignKey:DailyLogID"`
	WorkoutEntries       []WorkoutEntry `json:"workout_entries" gorm:"foreignKey:DailyLogID"`
	Sleep                float32        `json:"sleep"`
	TotalCalorieConsumed float32        `json:"total_calorie_consumed"`
	TotalCalorieTarget   float32        `json:"total_calorie_target"`
	CalorieBurned        int32          `json:"calorie_burned"`
	CaloriesBurnedTarget int32          `json:"calories_burn_target"`
	Date                 time.Time      `json:"date"`
	WaterTarget          int32          `json:"water_target"`
	CalorieBudgets       CalorieBudget  `json:"calorie_budgets"`
	CaloriesPerMeal      CalorieBudget  `json:"calories_per_meal"`
	NutritionTarget      Nutrition      `json:"nutrition_target"`
	NutritionConsumed    Nutrition      `json:"nutrition_consumed"`
}

// NewDailyLog initializes a new DailyLog with default values
func NewDailyLog(userID uint) *DailyLog {
	// TODO: implement this function
	// get user from db and update daily log
	// write funtion to greate calorie budget and nutrition target
	return &DailyLog{
		UserID:               userID,
		WaterEntries:         0,
		FoodEntries:          []FoodEntry{},
		WorkoutEntries:       []WorkoutEntry{},
		Sleep:                0,
		TotalCalorieConsumed: 0, // update whenever foodEntry is added, updated or removed
		TotalCalorieTarget:   0, // get from user
		CalorieBurned:        0, // update whenever workoutEntry is added, updated or removed
		CaloriesBurnedTarget: 0, // get from user
		Date:                 time.Now(),
		WaterTarget:          0,               // get from user
		CalorieBudgets:       CalorieBudget{}, // write a function to calculate this
		CaloriesPerMeal:      CalorieBudget{}, // update whenever foodEntry is added, updated or removed
		NutritionTarget:      Nutrition{},     // write a function to calculate this
		NutritionConsumed:    Nutrition{},     // update whenever foodEntry is added, updated or removed
	}
}
