package repo

import "github.com/fouadkada/k8s-tutorial/movies/internal/model"

var allMovies = model.Movies{
	model.Movie{
		Name: "Movie title 1",
		Year: "2009",
	},
	model.Movie{
		Name: "Movie title 2",
		Year: "2010",
	},
	model.Movie{
		Name: "Movie title 3",
		Year: "2011",
	},
	model.Movie{
		Name: "Movie title 4",
		Year: "2012",
	},
	model.Movie{
		Name: "Movie title 5",
		Year: "2013",
	},
}

type MoviesRepo struct {
}

func (r MoviesRepo) GetAllMovies() model.Movies {
	return allMovies
}
