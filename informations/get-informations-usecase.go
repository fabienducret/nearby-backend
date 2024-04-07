package informations

import (
	"nearby/models"
)

type WeatherRepository interface {
	Search(city string) (models.Weather, error)
}

func GetInformationsUseCaseFactory(weatherRepository WeatherRepository) GetInformationsUseCase {
	return func(city string) (models.Informations, error) {
		weather, _ := weatherRepository.Search(city)

		return models.Informations{Weather: weather}, nil
	}
}
