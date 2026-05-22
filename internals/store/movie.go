package store

import (
	"github.com/chirag-3/crud-api-in-memory/internals/models"
)

var movies []models.Movie

func InitializeStore() {
	movies = make([]models.Movie, 0)
}

func AddMovie(m models.Movie) {
	movies = append(movies, m)
}

func ReturnMovie(id int) models.Movie {
	for _, movie := range movies {
		if movie.Id == id {
			return movie
		}
	}
	return models.Movie{Id: 0, Title: "", Director: ""}
}
