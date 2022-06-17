package http

import (
	"github.com/fouadkada/k8s-tutorial/weather/internal/service"
	"net/http"
)

type handler struct {
	weatherService service.Weather
}

func (h *handler) getWeather(w http.ResponseWriter, r *http.Request) {
	result := h.weatherService.GetWeather()
	respond(w, r, result)
}
