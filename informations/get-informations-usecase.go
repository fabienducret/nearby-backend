package informations

import (
	"log"
	"nearby/models"
	"sync"
)

type WeatherRepository func(city string) (models.Weather, error)

type NewsRepository func(city string) ([]models.News, error)

func GetInformationsUseCaseFactory(weatherFor WeatherRepository, newsFor NewsRepository) GetInformationsUseCase {
	return func(city string) models.Informations {
		infos := models.Informations{}

		fetchWeather := func(city string) {
			weather, err := weatherFor(city)
			if err != nil {
				log.Println(err)
			}

			infos.Weather = weather
		}

		fetchNews := func(city string) {
			news, err := newsFor(city)
			if err != nil {
				log.Println(err)
			}

			infos.News = news
		}

		withConcurrency(city, []func(city string){fetchWeather, fetchNews})

		return infos
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
