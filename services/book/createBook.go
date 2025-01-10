package services_book

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	models_book "github.com/render-examples/go-gin-web-server/models/book"
)

// CreateBook godoc
// @Summary create book by json body
// @Description create book by json body as models_book.BookRequest then return as json message
// @Tags books
// @Accept json
// @Produce json
// @Param title query string true "string of book title" default(Some book title)
// @Param author query string true "string of book author" default(my self)
// @Param categories query []int true "category fill as number between 1-4" "Exmaple from bookCategories model" collectionFormat(multi)
// @Param is_borrow query bool true "constant of borrow status true is mean borrowed" default(false)
// @Param borrowed_count query int true "borrowed counter" default(0)
// @Success      200  {object}  handler.SuccessResponseModel
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /create/book/ [post]
func (m *BookService) CreateBook(c *gin.Context) {
	var body models_book.BookRequest

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	res, err := m.collection.InsertOne(context.TODO(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	if res.Acknowledged == false {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{Error: "err invalid book json struct"})
		return
	}

	c.JSON(http.StatusOK, handler.SuccessResponseModel{Message: "Book :" + body.Title + " was created"})
}
