package handler_lotto

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	struct_lotto "github.com/render-examples/go-gin-web-server/struct/lotto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (m *LottoService) GetAllLottoHandler(c *gin.Context) {
	var lotties []struct_lotto.Lotto
	opts := options.Find().SetSort(bson.M{"date": -1}) // DESC
	cursor, err := m.collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: err.Error(),
		})
		return
	}

	if err = cursor.All(context.TODO(), &lotties); err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   lotties,
	})
}
