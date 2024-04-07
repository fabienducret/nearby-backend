package weather

import (
	"fmt"
	"io"
	"nearby/models"
	"net/http"
)

type WeatherRepository struct {
	ApiKey string
}

func NewWeatherRepository(apiKey string) *WeatherRepository {
	return &WeatherRepository{apiKey}
}

func (r *WeatherRepository) Search(city string) (models.Weather, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&units=metric&q=%s", r.ApiKey, city), nil)
	if err != nil {
		return models.Weather{}, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return models.Weather{}, err
	}

	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Weather{}, err
	}

	return Parse(respBytes), nil
}
