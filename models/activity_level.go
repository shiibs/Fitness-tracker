package models

type ActivityLevel string

const (
	Sedentary        ActivityLevel = "sedentary"
	LightlyActive    ActivityLevel = "lightlyActive"
	ModeratelyActive ActivityLevel = "moderatelyActive"
	VeryActive       ActivityLevel = "veryActive"
	SuperActive      ActivityLevel = "superActive"
)
