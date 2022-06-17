package http

import (
	"github.com/fouadkada/k8s-tutorial/aggregator/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type Server struct {
	aggregatorService service.Aggregator
}

func NewAggregatorServer(aggregatorService service.Aggregator) Server {
	return Server{
		aggregatorService: aggregatorService,
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
		aggregatorService: s.aggregatorService,
	}

	r.Get("/all", rest.getAggregated)

	return http.ListenAndServe(":3000", r)
}
