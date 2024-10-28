package models

import "time"

type WeightEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Weight    float32   `json:"weight"`
	Timestamp time.Time `json:"timestamp"`
}
