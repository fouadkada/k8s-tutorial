package http

import (
	"github.com/fouadkada/k8s-tutorial/movies/internal/service"
	"net/http"
)

type handler struct {
	moviesService service.Movie
}

func (h *handler) getAllMovies(w http.ResponseWriter, r *http.Request) {
	result := h.moviesService.GetAllMovies()
	respond(w, r, result)
}
