package models

type Recipe struct {
	ID          uint               `json:"id" gorm:"primaryKey"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Ingredients []RecipeIngredient `json:"ingredients" gorm:"foreignKey:RecipeID"`
	Steps       []RecipeStep       `json:"steps" gorm:"foreignKey:RecipeID"`
	Nutrition   Nutrition          `json:"nutrition"`
}

type RecipeIngredient struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	RecipeID uint    `json:"recipe_id"`
	FoodID   uint    `json:"food_id"`
	Quantity float32 `json:"quantity"`
	Unit     Unit    `json:"unit"`
}

type RecipeStep struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	RecipeID uint   `json:"recipe_id"`
	Step     string `json:"step"`
	Order    int    `json:"order"`
}
