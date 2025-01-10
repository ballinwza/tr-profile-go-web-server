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

// BorrowBook godoc
// @Summary changing isBorrow status
// @Description changing isBorrow status true is mean borrowed this route will increase borrowed_count plus 1 then return json message
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "string id of book" "Example : 67809c673105c8f80c845b8e"
// @Success      200  {object}  handler.SuccessResponseModel
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /borrow/book/{id} [put]
func (m *BookService) BorrowBook(c *gin.Context) {
	var result models_book.BookRequest
	paramId := c.Param("id")
	objId, err := bson.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	filter := bson.M{"_id": objId}
	update := bson.A{
		bson.M{
			"$set": bson.M{"is_borrow": bson.M{"$not": bson.A{"$is_borrow"}}},
		},
		bson.M{
			"$set": bson.M{
				"borrowed_count": bson.M{
					"$cond": bson.M{
						"if":   "$is_borrow",
						"then": bson.M{"$add": bson.A{"$borrowed_count", 1}},
						"else": "$borrowed_count",
					},
				},
			},
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = m.collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	if result.IsBorrow == true {
		c.JSON(http.StatusOK, handler.SuccessResponseModel{
			Message: "Book : " + result.Title + " was borrowed",
		})
		return
	}
	c.JSON(http.StatusOK, handler.SuccessResponseModel{
		Message: "Book : " + result.Title + " was returned",
	})

}
