package services_book

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// DeleteBook godoc
// @Summary delete book by id
// @Description delete book by id then return as message
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "string id of book" "Example : 67809c673105c8f80c845b8e"
// @Success      200  {object}  handler.SuccessResponseModel
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /delete/book/{id} [delete]
func (m *BookService) DeleteBook(c *gin.Context) {
	paramId := c.Param("id")
	objId, err := bson.ObjectIDFromHex(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{Error: err.Error()})
		return
	}

	filter := bson.M{"_id": objId}
	res, err := m.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{Error: err.Error()})
		return
	}

	if res.Acknowledged != true {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{Error: "error invalid id"})
		return
	}

	c.JSON(http.StatusOK, handler.SuccessResponseModel{Message: "Book ID : " + paramId + " was deleted"})
}
