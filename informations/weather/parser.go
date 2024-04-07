package weather

import (
	"encoding/json"
	"nearby/models"
)

func Parse(from []byte) models.Weather {
	var reply struct {
		Main struct {
			Temp float32 `json:"temp"`
		} `json:"main"`
	}

	json.Unmarshal(from, &reply)

	return models.Weather{
		Temperature: reply.Main.Temp,
	}
}
