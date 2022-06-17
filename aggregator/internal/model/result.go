package model

import (
	"net/http"
)

type Movie struct {
	Name string
	Year string
}

func (m Movie) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Weather struct {
	City      string
	Condition string
}

func (w2 Weather) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Result struct {
	Movies  []Movie
	Weather []Weather
}

func (r2 Result) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
