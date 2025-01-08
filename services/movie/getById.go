package services_movie

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	"github.com/render-examples/go-gin-web-server/models/models_movie"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetMovieById godoc
// @Summary Response movie by id
// @Description get movie by id param then return movie detail as json
// @Tags movies
// @Accept json
// @Produce json
// @Param id path string true "id" default(2025)
// @Success      200  {object}  responseMovie
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /movie/{id} [get]
func (m *MovieService) GetMovieById(c *gin.Context) {
	var result models_movie.MovieRes
	paramId := c.Param("id")
	objectId, err := bson.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	filter := bson.M{"id": objectId}
	err = m.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responseMovie{Data: result})
}
