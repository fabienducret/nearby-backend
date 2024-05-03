package informations_test

import (
	"fmt"
	"nearby/informations"
	"nearby/models"
	"reflect"
	"testing"
)

type spyLogger struct {
	hasBeenCalled bool
}

func (l *spyLogger) Println(v ...any) {
	l.hasBeenCalled = true
}

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
		logger := &spyLogger{}
		getInfos := informations.InitGetInformationsUseCase(stubWeatherRepository, subNewsRepository, logger)

		informations := getInfos("Paris")

		assertDeepEqual(t, informations, want)
		assertEqual(t, logger.hasBeenCalled, false)
	})

	t.Run("get empty news because of error", func(t *testing.T) {
		want := models.Informations{
			Weather: models.Weather{Temperature: 15.5},
			News:    []models.News{},
		}
		logger := &spyLogger{}
		subNewsRepository := func(city string) ([]models.News, error) {
			return []models.News{}, fmt.Errorf("error in news")
		}

		getInfos := informations.InitGetInformationsUseCase(stubWeatherRepository, subNewsRepository, logger)

		informations := getInfos("Paris")

		assertDeepEqual(t, informations, want)
		assertEqual(t, logger.hasBeenCalled, true)
	})
}

func assertDeepEqual(t *testing.T, got, want models.Informations) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error in assert deep equal, got %+v, want %+v", got, want)
	}
}

func assertEqual[T bool](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("Error in assert equal, got %+v, want %+v", got, want)
	}
}
