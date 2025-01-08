package services_movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	"github.com/render-examples/go-gin-web-server/models/models_movie"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetAllMovie godoc
// @Summary Response movie list by filter
// @Description get movie list by filter model then return movie list as json
// @Tags movies
// @Accept json
// @Produce json
// @Success      200  {object}  responseMovieList
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /movie/list [get]
func (m *MovieService) GetAllMovie(c *gin.Context) {
	var result []models_movie.MovieRes
	filter := bson.M{}
	cursor, err := m.collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responseMovieList{Data: result})

}
