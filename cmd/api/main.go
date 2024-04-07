package main

import (
	"nearby/health"
	"nearby/informations"
	"nearby/informations/weather"
	"nearby/server"
	"os"
)

func main() {
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	weatherRepository := weather.NewWeatherRepository(weatherApiKey)
	getInformationsUseCase := informations.GetInformationsUseCaseFactory(weatherRepository)

	c := server.Controllers{
		Health:       health.Controller,
		Informations: informations.ControllerFactory(getInformationsUseCase),
	}

	s := server.New(c)
	s.Run("8080")
}
