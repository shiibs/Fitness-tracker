package service

import (
	"github.com/shiibs/fitness-app/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{db}
}

func (s *UserService) UpdateUser(userId uint, user *models.User) (*models.User, error) {
	var existingUser models.User
	err := s.db.First(&existingUser, userId).Error

	if err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.ImageURL = user.ImageURL
	existingUser.Age = user.Age
	existingUser.Height = user.Height
	existingUser.CurrentWeight = user.CurrentWeight
	existingUser.Goal = user.Goal
	existingUser.TargetWeight = user.TargetWeight
	existingUser.ActivityLevel = user.ActivityLevel
	existingUser.Gender = user.Gender
	existingUser.FoodPreference = user.FoodPreference
	existingUser.BMR = calculateBMR(existingUser.Gender, existingUser.CurrentWeight, float32(existingUser.Height), float32(existingUser.Age))
	existingUser.TDEE = calculateTDEE(existingUser.BMR, existingUser.ActivityLevel)
	existingUser.DailyCaloricIntake = existingUser.TDEE
	if existingUser.Goal == models.WeightLose {
		existingUser.DailyCaloricIntake -= 300
	}
	existingUser.WaterTarget = calculateWaterTarget(existingUser.BMR)
	existingUser.CalorieBurnTarget = int32(calculateCalorieBurnTarget(existingUser.ActivityLevel))

	if err := s.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func calculateBMR(gender models.Gender, weight, height, age float32) float32 {
	if gender == models.Male {
		return 88.362 + (13.397 * weight) + (4.799 * height) - (5.677 * age)
	}
	return 447.593 + (9.247 * weight) + (3.098 * height) - (4.330 * age)
}

func calculateTDEE(bmr float32, activityLevel models.ActivityLevel) float32 {
	var multiplier float32
	switch activityLevel {
	case models.Sedentary:
		multiplier = 1.2
	case models.LightlyActive:
		multiplier = 1.375
	case models.ModeratelyActive:
		multiplier = 1.55
	case models.VeryActive:
		multiplier = 1.725
	case models.SuperActive:
		multiplier = 1.9
	default:
		multiplier = 1.2
	}

	return bmr * multiplier
}

func calculateCalorieBurnTarget(activityLevel models.ActivityLevel) float32 {
	switch activityLevel {
	case models.Sedentary:
		return 250
	case models.LightlyActive:
		return 300
	default:
		return 400
	}
}

func calculateWaterTarget(bmr float32) int32 {
	waterInLiters := bmr * 0.03
	waterInGlass := waterInLiters / 0.25
	return int32(waterInGlass)
}
