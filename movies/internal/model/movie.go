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

type Movies []Movie

func (m Movies) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
