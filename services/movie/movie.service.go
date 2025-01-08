package services_movie

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/connecter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IMovieService interface {
	CreateMovie(*gin.Context)
	GetAllMovie(c *gin.Context)
	GetMovieById(*gin.Context)
}

type MovieService struct {
	collection *mongo.Collection
}

func SetupMovieService() *MovieService {
	return &MovieService{
		collection: connecter.ConnectWithMongo("tr-profile-go-movie", "my-favorite-movies"),
	}
}
