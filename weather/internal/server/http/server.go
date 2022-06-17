package http

import (
	"github.com/fouadkada/k8s-tutorial/weather/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type Server struct {
	weatherService service.Weather
}

func NewWeatherServer(weatherService service.Weather) Server {
	return Server{
		weatherService: weatherService,
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
		weatherService: s.weatherService,
	}

	r.Get("/weather", rest.getWeather)

	return http.ListenAndServe(":3000", r)
}
