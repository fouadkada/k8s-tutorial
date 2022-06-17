package http

import (
	"github.com/fouadkada/k8s-tutorial/aggregator/internal/service"
	"net/http"
)

type handler struct {
	aggregatorService service.Aggregator
}

func (h *handler) getAggregated(w http.ResponseWriter, r *http.Request) {
	result := h.aggregatorService.Aggregate()
	respond(w, r, result)
}
