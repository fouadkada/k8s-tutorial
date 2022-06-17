package main

import (
	"github.com/fouadkada/k8s-tutorial/movies/internal/repo"
	"github.com/fouadkada/k8s-tutorial/movies/internal/server/http"
	"github.com/fouadkada/k8s-tutorial/movies/internal/service"
	"log"
)

func main() {
	repository := repo.MoviesRepo{}
	srvc := service.NewMoviesService(repository)
	server := http.NewMoviesServer(srvc)
	log.Fatalln(server.StartServer())
}
