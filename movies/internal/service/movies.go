package service

import (
	"github.com/fouadkada/k8s-tutorial/movies/internal/model"
	"github.com/fouadkada/k8s-tutorial/movies/internal/repo"
)

type Movie struct {
	moviesRepo repo.MoviesRepo
}

func NewMoviesService(moviesRepo repo.MoviesRepo) Movie {
	return Movie{
		moviesRepo: moviesRepo,
	}
}

func (s Movie) GetAllMovies() model.Movies {
	return s.moviesRepo.GetAllMovies()
}
