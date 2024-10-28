package models

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type FoodPreference string

const (
	Vegan         FoodPreference = "vegan"
	Vegetarian    FoodPreference = "vegetarian"
	OvoVegetarian FoodPreference = "ovoVegetarian"
	NonVegetarian FoodPreference = "nonVegetarian"
)

type User struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	GoogleID          string         `json:"google_id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	ImageURL          string         `json:"image_url"`
	Age               int            `json:"age"`
	Height            int            `json:"height"` // in cm
	CurrentWeight     float32        `json:"current_weight"`
	Goal              Goal           `json:"goal"`
	TargetWeight      float32        `json:"target_weight"`
	ActivityLevel     ActivityLevel  `json:"activity_level"`
	Gender            Gender         `json:"gender"`
	FoodPreference    FoodPreference `json:"food_preference"`
	BMR               float32        `json:"bmr"`
	TDEE              float32        `json:"tdee"`
	CalorieBurnTarget int32          `json:"calorie_burn_target"`
	WeightEntries     []WeightEntry  `json:"weight_entries" gorm:"foreignKey:UserID"` // Foreign key to WeightEntry
}
