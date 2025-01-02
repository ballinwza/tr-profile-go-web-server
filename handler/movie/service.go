package handler_movie

import (
	"github.com/render-examples/go-gin-web-server/connecter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MovieService struct {
	collection *mongo.Collection
}

func SetupMovieService() *MovieService {
	return &MovieService{
		collection: connecter.ConnectWithMongo("tr-profile-go-movie", "my-favorite-movies"),
	}
}
