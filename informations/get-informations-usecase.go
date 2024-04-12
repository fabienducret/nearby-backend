package informations

import (
	"fmt"
	"nearby/models"
)

type WeatherRepository func(city string) (models.Weather, error)

type NewsRepository func(city string) ([]models.News, error)

func GetInformationsUseCaseFactory(weatherFor WeatherRepository, newsFor NewsRepository) GetInformationsUseCase {
	fetchWeather := func(city string, result chan<- models.Weather) {
		weather, err := weatherFor(city)
		if err != nil {
			fmt.Println(err)
		}

		result <- weather
	}

	fetchNews := func(city string, result chan<- []models.News) {
		news, err := newsFor(city)
		if err != nil {
			fmt.Println(err)
		}

		result <- news
	}

	return func(city string) models.Informations {
		weather := make(chan models.Weather)
		news := make(chan []models.News)

		go fetchWeather(city, weather)
		go fetchNews(city, news)

		return models.Informations{
			Weather: <-weather,
			News:    <-news,
		}
	}
}
