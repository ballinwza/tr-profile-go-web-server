package services_lotto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	struct_lotto "github.com/render-examples/go-gin-web-server/struct/lotto"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Test struct {
	data struct_lotto.Lotto
}

type Err struct {
	Err  error
	Type uint64
	Meta any
}

type T struct {
	id string `example:"677accbf56ff1bc2a731b944"`
}

// GetLottoById godoc
// @Summary      Show an lotto
// @Description  get string by id then return lotto
// @Tags         lotties
// @Accept       json
// @Produce      json
// @Param        id   path      string true  "Exmaple : 677accbf56ff1bc2a731b944"
// @Success      200  {object}  Test
// @Failure      400  {object}  Err
// @Failure      404  {object}  map[string]error
// @Failure      500  {object}  map[string]error
// @Router       /lotto/{id} [get]
func (m *LottoService) GetLottoByIdHandler(c *gin.Context) {
	var lotto struct_lotto.Lotto
	param := c.Param("id")
	objectId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
			Meta: err.Error(),
		})
		return
	}

	err = m.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&lotto)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.Error{
			Err:  err,
			Type: http.StatusBadRequest,
			Meta: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": lotto,
	})
}
