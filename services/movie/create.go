package services_movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	"github.com/render-examples/go-gin-web-server/models/models_movie"
)

func (m *MovieService) CreateMovie(c *gin.Context) {
	var body models_movie.CreateMovieReq

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	_, err := m.collection.InsertOne(context.TODO(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, handler.SuccessResponseModel{Message: "Movie : " + body.Title + " was created"})
}
