package services_book

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	models_book "github.com/render-examples/go-gin-web-server/models/book"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetMostBorrowBook godoc
// @Summary find most popular book
// @Description find most popular book then return as json responseBook model
// @Tags books
// @Accept json
// @Produce json
// @Success      200  {object}  responseBook
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /book/popular [get]
func (m *BookService) GetMostBorrowBook(c *gin.Context) {
	var result []models_book.BookFilterResponse

	pipeline := mongo.Pipeline{
		{{Key: "$lookup", Value: bson.M{
			"from":         "category",
			"localField":   "categories",
			"foreignField": "id",
			"as":           "categoryDetail",
		},
		}},
		{{Key: "$unwind", Value: bson.M{"path": "$categoryDetail"}}},
		{{Key: "$group", Value: bson.M{
			"_id":            "$_id",
			"title":          bson.M{"$first": "$title"},
			"author":         bson.M{"$first": "$author"},
			"categories":     bson.M{"$push": "$categoryDetail.name"},
			"is_borrow":      bson.M{"$first": "$is_borrow"},
			"borrowed_count": bson.M{"$first": "$borrowed_count"},
		}}},
		{{Key: "$sort", Value: bson.M{"title": -1}}},
		{{Key: "$limit", Value: 1}},
	}

	cursor, err := m.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{Error: err.Error()})
		return
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{Error: err.Error()})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{Error: "error not found any book"})
		return
	}

	c.JSON(http.StatusOK, responseBook{
		Data: result[0],
	})
}
