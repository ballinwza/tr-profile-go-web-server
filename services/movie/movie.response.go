package services_movie

import "github.com/render-examples/go-gin-web-server/models/models_movie"

type responseMovie struct {
	Data models_movie.MovieRes `json:"data"`
}

type responseMovieList struct {
	Data []models_movie.MovieRes `json:"data"`
}
