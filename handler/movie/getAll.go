package handler_movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	struct_movie "github.com/render-examples/go-gin-web-server/struct/movie"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (m *MovieService) GetAllMovieHandler(c *gin.Context) {
	var result []struct_movie.MovieRes
	filter := bson.M{}
	cursor, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
		})
		return
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(handler.SuccessAndResponseJson(result))

}
