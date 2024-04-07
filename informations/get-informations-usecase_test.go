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

func TestGetInformationsUseCase(t *testing.T) {
	t.Run("get weather informations with success", func(t *testing.T) {
		want := models.Informations{Weather: models.Weather{Temperature: 15.5}}
		getInfos := informations.GetInformationsUseCaseFactory(stubWeatherRepository)

		informations, _ := getInfos("Paris")

		assertDeepEqual(t, informations, want)
	})
}

func assertDeepEqual(t *testing.T, got, want models.Informations) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error in assert deep equal, got %+v, want %+v", got, want)
	}
}
