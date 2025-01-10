package services_movie

import models_movie "github.com/render-examples/go-gin-web-server/models/movie"

type responseMovie struct {
	Data models_movie.MovieRes `json:"data"`
}

type responseMovieList struct {
	Data []models_movie.MovieRes `json:"data"`
}
