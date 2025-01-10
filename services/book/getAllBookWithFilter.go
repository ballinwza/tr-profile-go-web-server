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

// GetAllBookWithFilter godoc
// @Summary get book list by filter
// @Description get book list from title, author or categoyies then return as json message
// @Tags books
// @Accept json
// @Produce json
// @Param title query string false "string of book title" default(The Witcher Blood of Elves)
// @Param author query string false "string of author" default(Andrzej Sapkowski)
// @Param categories query []int false "integer of categories id between 1-4" "Exmaple from bookCategories model" collectionFormat(multi)
// @Success      200  {object}  responseBookList
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /book/list [get]
func (m *BookService) GetAllBookWithFilter(c *gin.Context) {
	var books []models_book.BookFilterResponse
	var req models_book.BookFilterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}
	filter := bson.M{}

	if req.Title != nil {
		filter["title"] = req.Title
	}
	if req.Author != nil {
		filter["author"] = req.Author
	}
	if req.Categories != nil {
		filter["categories"] = bson.M{"$in": req.Categories}
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
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
		{{Key: "$sort", Value: bson.M{"title": 1}}},
	}

	cursor, err := m.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	if err = cursor.All(context.TODO(), &books); err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responseBookList{
		Data: books,
	})
}
