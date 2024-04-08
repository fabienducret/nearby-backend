package informations

import (
	"nearby/models"
	"sync"
)

type WeatherRepository func(city string) (models.Weather, error)

type NewsRepository func(city string) ([]models.News, error)

func GetInformationsUseCaseFactory(weatherFor WeatherRepository, newsFor NewsRepository) GetInformationsUseCase {
	return func(city string) (models.Informations, error) {
		infos := models.Informations{}

		fetchWeather := func(city string) {
			weather, _ := weatherFor(city)
			infos.Weather = weather
		}

		fetchNews := func(city string) {
			news, _ := newsFor(city)
			infos.News = news
		}

		withConcurrency(city, []func(city string){fetchWeather, fetchNews})

		return infos, nil
	}
}

func withConcurrency(city string, functions []func(city string)) {
	wg := sync.WaitGroup{}

	for _, f := range functions {
		wg.Add(1)

		go func(city string) {
			defer wg.Done()
			f(city)
		}(city)
	}

	wg.Wait()
}
