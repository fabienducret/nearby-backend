package weather

import (
	"fmt"
	"io"
	"nearby/models"
	"net/http"
	"time"
)

func InitWeatherRepository(apiKey string) func(city string) (models.Weather, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?appid=%s&units=metric", apiKey)

	return func(city string) (models.Weather, error) {
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s&q=%s", url, city), nil)
		if err != nil {
			return models.Weather{}, err
		}

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
}
