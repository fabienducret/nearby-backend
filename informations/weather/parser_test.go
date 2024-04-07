package weather_test

import (
	"nearby/informations/weather"
	"nearby/models"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Parse weather with success", func(t *testing.T) {
		toParse := []byte(`{"coord":{"lon":4,"lat":46},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"base":"stations","main":{"temp":20.94,"feels_like":20.66,"temp_min":19.97,"temp_max":20.95,"pressure":1013,"humidity":60,"sea_level":1013,"grnd_level":954},"visibility":10000,"wind":{"speed":3.71,"deg":154,"gust":9.63},"clouds":{"all":100},"dt":1712495477,"sys":{"type":2,"id":2038399,"country":"FR","sunrise":1712466708,"sunset":1712513992},"timezone":7200,"id":2983360,"name":"Arrondissement de Roanne","cod":200}`)
		want := models.Weather{
			Temperature:          20.94,
			FeelsLikeTemperature: 20.66,
			Pressure:             1013,
			Humidity:             60,
		}

		weather := weather.Parse(toParse)

		assertDeepEqual(t, weather, want)
	})

}

func assertDeepEqual(t *testing.T, got, want models.Weather) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Error in assert deep equal, got %+v, want %+v", got, want)
	}
}
