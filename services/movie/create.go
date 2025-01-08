package services_movie

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/render-examples/go-gin-web-server/handler"
	struct_movie "github.com/render-examples/go-gin-web-server/struct/movie"
)

func (m *MovieService) CreateMovieHandler(c *gin.Context) {
	var body struct_movie.CreateMovieReq
	newId := uuid.New().String()

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
			Meta: map[string]string{
				"message": "Invalid input movie",
			},
		})
		return
	}

	if body.Uuid != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  fmt.Errorf("do not input ID"),
			Type: http.StatusBadRequest,
			Meta: map[string]string{
				"message": "Do not input ID",
			},
		})
		return
	}

	body.Uuid = &newId
	res, err := m.collection.InsertOne(context.TODO(), body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(handler.SuccessAndResponseJson(res.InsertedID))
}
