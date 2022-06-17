package main

import (
	"github.com/fouadkada/k8s-tutorial/weather/internal/repo"
	"github.com/fouadkada/k8s-tutorial/weather/internal/server/http"
	"github.com/fouadkada/k8s-tutorial/weather/internal/service"
	"log"
)

func main() {
	repository := repo.WeatherRepo{}
	srvc := service.NewWeatherService(repository)
	server := http.NewWeatherServer(srvc)
	log.Fatalln(server.StartServer())
}
