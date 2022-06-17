package http

import (
	"github.com/fouadkada/k8s-tutorial/movies/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type Server struct {
	moviesService service.Movie
}

func NewMoviesServer(moviesService service.Movie) Server {
	return Server{
		moviesService: moviesService,
	}
}

func (s *Server) StartServer() error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	rest := handler{
		moviesService: s.moviesService,
	}

	r.Get("/movies", rest.getAllMovies)

	return http.ListenAndServe(":3001", r)
}
