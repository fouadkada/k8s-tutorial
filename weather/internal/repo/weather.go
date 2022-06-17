package repo

import "github.com/fouadkada/k8s-tutorial/weather/internal/model"

var allCities = []model.Weather{
	model.Weather{
		City:      "Beirut",
		Condition: "sunny",
	},
	model.Weather{
		City:      "Vienna",
		Condition: "grey",
	},
	model.Weather{
		City:      "Dubai",
		Condition: "sunny",
	},
	model.Weather{
		City:      "New York",
		Condition: "foggy",
	},
}

type WeatherRepo struct {
}

func (r WeatherRepo) GetWeather() model.WeatherCollection {
	return allCities
}
