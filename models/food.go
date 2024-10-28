package models

type Food struct {
	ID            uint               `json:"id" gorm:"primaryKey"`
	Name          string             `json:"name"`
	UnitNutrients map[Unit]Nutrition `json:"unit_nutrients"` // A map of units to their corresponding nutritional values
}
