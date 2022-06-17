package model

import "net/http"

type Weather struct {
	City      string
	Condition string
}

func (w2 Weather) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type WeatherCollection []Weather

func (w2 WeatherCollection) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
