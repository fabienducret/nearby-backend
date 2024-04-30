package config

import "os"

type Config struct {
	WeatherApiKey string
}

func MustInit() *Config {
	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	if weatherApiKey == "" {
		panic("missing WEATHER_API_KEY")
	}

	return &Config{
		WeatherApiKey: weatherApiKey,
	}
}
