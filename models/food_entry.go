package models

type FoodEntry struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	UserID     uint     `json:"user_id"`
	DailyLogID uint     `json:"daily_log_id" gorm:"foreignKey:DailyLogID"`
	FoodID     uint     `json:"food_id"`
	Quantity   float32  `json:"quantity"`
	Unit       Unit     `json:"unit"`
	MealType   MealType `json:"meal_type"`
}
