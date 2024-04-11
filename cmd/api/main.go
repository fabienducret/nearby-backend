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
	if weatherApiKey == "" {
		panic("missing WEATHER_API_KEY")
	}

	c := controllers(weatherApiKey)
	s := server.New(c)
	s.Run("8080")
}

func controllers(weatherApiKey string) server.Controllers {
	weatherRepository := weather.WeatherRepositoryFactory(weatherApiKey)
	newsRepository := news.NewsRepositoryFactory()
	getInformationsUseCase := informations.GetInformationsUseCaseFactory(weatherRepository, newsRepository)

	return server.Controllers{
		Health:       health.Controller,
		Informations: informations.ControllerFactory(getInformationsUseCase),
	}
}
