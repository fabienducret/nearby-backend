package main

import (
	"nearby/health"
	"nearby/informations"
	"nearby/informations/news"
	"nearby/informations/weather"
	"nearby/server"
	"os"
)

func main() {
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	weatherRepository := weather.WeatherRepositoryFactory(weatherApiKey)
	newsRepository := news.NewsRepositoryFactory()
	getInformationsUseCase := informations.GetInformationsUseCaseFactory(weatherRepository, newsRepository)

	c := server.Controllers{
		Health:       health.Controller,
		Informations: informations.ControllerFactory(getInformationsUseCase),
	}

	s := server.New(c)
	s.Run("8080")
}
