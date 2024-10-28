package models

type Goal string

const (
	WeightLose Goal = "lose"
	Maintain   Goal = "maintain"
)

type ActivityLevel string

const (
	Sedentary        ActivityLevel = "sedentary"
	LightlyActive    ActivityLevel = "lightlyActive"
	ModeratelyActive ActivityLevel = "moderatelyActive"
	VeryActive       ActivityLevel = "veryActive"
	SuperActive      ActivityLevel = "superActive"
)

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
	ID             uint
	GoogleID       string
	Name           string
	Email          string
	ImageURL       string
	Age            int
	Height         int // in cm
	CurrentWeight  float32
	Goal           Goal
	TargetWeight   float32
	ActivityLevel  ActivityLevel
	Gender         Gender
	FoodPreference FoodPreference
	BMR            float32
	TDEE           float32
}
