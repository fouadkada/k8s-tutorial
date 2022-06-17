package service

import (
	"github.com/fouadkada/k8s-tutorial/weather/internal/model"
	"github.com/fouadkada/k8s-tutorial/weather/internal/repo"
)

type Weather struct {
	weatherRepo repo.WeatherRepo
}

func NewWeatherService(weatherRepo repo.WeatherRepo) Weather {
	return Weather{
		weatherRepo: weatherRepo,
	}
}

func (s Weather) GetWeather() model.WeatherCollection {
	return s.weatherRepo.GetWeather()
}
