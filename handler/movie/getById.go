package handler_movie

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	struct_movie "github.com/render-examples/go-gin-web-server/struct"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (m *MovieService) GetMovieByIdHandler(c *gin.Context) {
	var result struct_movie.Movie
	var param struct_movie.BindMovieById
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
		})
		return
	}

	conIdToNum, err := strconv.Atoi(param.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
			Meta: map[string]string{
				"message": "id must be number",
			},
		})
		return
	}

	filter := bson.M{"_id": conIdToNum}
	err = m.collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(result)
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

	c.JSON(http.StatusOK, gin.H{"data": result})
}
