package main

import (
	"github.com/fouadkada/k8s-tutorial/aggregator/internal/server/http"
	"github.com/fouadkada/k8s-tutorial/aggregator/internal/service"
	"log"
)

func main() {
	srvc := service.NewAggregatorService("http://movies:3000/movies", "http://weather:3000/weather")
	server := http.NewAggregatorServer(srvc)
	log.Fatalln(server.StartServer())
}
