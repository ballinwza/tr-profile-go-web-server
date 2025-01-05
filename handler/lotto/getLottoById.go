package handler_lotto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	struct_lotto "github.com/render-examples/go-gin-web-server/struct/lotto"

	"go.mongodb.org/mongo-driver/v2/bson"
)

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
		"status": http.StatusOK,
		"data":   lotto,
	})
}
