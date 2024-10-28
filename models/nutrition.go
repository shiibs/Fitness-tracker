package models

type Nutrition struct {
	Calories float32 `json:"calories"`
	Protein  float32 `json:"protein"`
	Carbs    float32 `json:"carbs"`
	Fiber    float32 `json:"fiber"`
	Fat      float32 `json:"fat"`
}
