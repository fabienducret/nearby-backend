package models

type Weather struct {
	Temperature          float32 `json:"temperature"`
	FeelsLikeTemperature float32 `json:"feels_like_temperature"`
	Pressure             float32 `json:"pressure"`
	Humidity             int     `json:"humidity"`
}
