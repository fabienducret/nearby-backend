package informations_test

import (
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
		{Title: "fake news", Description: "this is a fake news"},
	}, nil
}

func TestGetInformationsUseCase(t *testing.T) {
	t.Run("get weather informations with success", func(t *testing.T) {
		want := models.Informations{
			Weather: models.Weather{Temperature: 15.5},
			News:    []models.News{{Title: "fake news", Description: "this is a fake news"}},
		}
		getInfos := informations.GetInformationsUseCaseFactory(stubWeatherRepository, subNewsRepository)

		informations, _ := getInfos("Paris")

		assertDeepEqual(t, informations, want)
	})
}

func assertDeepEqual(t *testing.T, got, want models.Informations) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error in assert deep equal, got %+v, want %+v", got, want)
	}
}
