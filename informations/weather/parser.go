package weather

import (
	"encoding/json"
	"nearby/models"
)

func Parse(from []byte) models.Weather {
	var reply struct {
		Main struct {
			Temp          float32 `json:"temp"`
			FeelsLikeTemp float32 `json:"feels_like"`
			Pressure      float32 `json:"pressure"`
			Humidity      int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}

	json.Unmarshal(from, &reply)

	return models.Weather{
		Temperature:          reply.Main.Temp,
		FeelsLikeTemperature: reply.Main.FeelsLikeTemp,
		Pressure:             reply.Main.Pressure,
		Humidity:             reply.Main.Humidity,
		Description:          reply.Weather[0].Description,
	}
}
