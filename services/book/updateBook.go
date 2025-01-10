package services_book

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	models_book "github.com/render-examples/go-gin-web-server/models/book"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UpdateBook godoc
// @Summary update book from param id
// @Description update book from param id from _id then return as json message
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Example : 678139d4e150ae9affc077bf"
// @Param title query string true "string of book title" default(Some book title)
// @Param author query string true "string of book author" default(my self)
// @Param categories query []int true "category fill as number between 1-4" "Exmaple from bookCategories model" collectionFormat(multi)
// @Param is_borrow query bool true "constant of borrow status true is mean borrowed" default(false)
// @Param borrowed_count query int true "borrowed counter" default(0)
// @Success      200  {object}  handler.SuccessResponseModel
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /update/book/{id} [put]
func (m *BookService) UpdateBook(c *gin.Context) {
	var body models_book.BookUpdateRequest
	var result models_book.BookDefaultResponse
	paramId := c.Param("id")
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	objId, err := bson.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	filter := bson.M{"_id": objId}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = m.collection.FindOneAndUpdate(context.TODO(), filter, bson.M{
		"$set": body,
	}, opts).Decode(&result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, handler.SuccessResponseModel{
		Message: "Book : " + result.Title + " was updated",
	})

}
