package informations_test

import (
	"fmt"
	"nearby/informations"
	"nearby/models"
	"reflect"
	"testing"
)

func stubWeatherRepository(city string) (models.Weather, error) {
	return models.Weather{
		Temperature: 15.5,
	}, nil
}

func subNewsRepository(city string) ([]models.News, error) {
	return []models.News{
		{Title: "fake news"},
	}, nil
}

func TestGetInformationsUseCase(t *testing.T) {
	t.Run("get informations with success", func(t *testing.T) {
		want := models.Informations{
			Weather: models.Weather{Temperature: 15.5},
			News:    []models.News{{Title: "fake news"}},
		}
		getInfos := informations.GetInformationsUseCaseFactory(stubWeatherRepository, subNewsRepository)

		informations := getInfos("Paris")

		assertDeepEqual(t, informations, want)
	})

	t.Run("get empty news because of error", func(t *testing.T) {
		want := models.Informations{
			Weather: models.Weather{Temperature: 15.5},
			News:    []models.News{},
		}

		subNewsRepository := func(city string) ([]models.News, error) {
			return []models.News{}, fmt.Errorf("error in news")
		}

		getInfos := informations.GetInformationsUseCaseFactory(stubWeatherRepository, subNewsRepository)

		informations := getInfos("Paris")

		assertDeepEqual(t, informations, want)
	})
}

func assertDeepEqual(t *testing.T, got, want models.Informations) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error in assert deep equal, got %+v, want %+v", got, want)
	}
}
