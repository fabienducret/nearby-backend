package main

import (
	"log"
	"nearby/config"
	"nearby/health"
	"nearby/informations"
	"nearby/informations/news"
	"nearby/informations/weather"
	"nearby/server"
)

func main() {
	config := config.MustInit()

	c := controllers(config)
	s := server.New(c)
	s.Run("8080")
}

func controllers(config *config.Config) server.Controllers {
	weatherRepository := weather.InitWeatherRepository(config.WeatherApiKey)
	newsRepository := news.InitNewsRepository()
	logger := log.Default()
	getInformationsUseCase := informations.InitGetInformationsUseCase(weatherRepository, newsRepository, logger)

	return server.Controllers{
		Health:       health.Controller,
		Informations: informations.InitController(getInformationsUseCase),
	}
}
