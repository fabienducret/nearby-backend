package informations

import (
	"nearby/models"
)

type WeatherRepository func(city string) (models.Weather, error)

func GetInformationsUseCaseFactory(weatherFor WeatherRepository) GetInformationsUseCase {
	return func(city string) (models.Informations, error) {
		weather, _ := weatherFor(city)

		return models.Informations{Weather: weather}, nil
	}
}
