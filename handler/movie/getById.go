package handler_movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	struct_movie "github.com/render-examples/go-gin-web-server/struct/movie"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (m *MovieService) GetMovieByIdHandler(c *gin.Context) {
	var result struct_movie.MovieRes
	var param struct_movie.BindMovieById
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
		})
		return
	}

	filter := bson.M{"id": param.Id}
	err := m.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
			Meta: map[string]string{
				"message": mongo.ErrNoDocuments.Error(),
			},
		})
		return
	}

	c.JSON(handler.SuccessAndResponseJson(result))
}
