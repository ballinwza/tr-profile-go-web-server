package services_lotto

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	struct_lotto "github.com/render-examples/go-gin-web-server/models/lotto"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// GetLottoById godoc
// @Summary      Show an lotto
// @Description  get string by id then return lotto
// @Tags         lotties
// @Accept       json
// @Produce      json
// @Param        id   path      string true  "Exmaple : 677accbf56ff1bc2a731b944"
// @Success      200  {object}  responseLotto
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /lotto/{id} [get]
func (m *LottoService) GetLottoById(c *gin.Context) {
	var lotto struct_lotto.Lotto
	param := c.Param("id")
	objectId, err := bson.ObjectIDFromHex(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	err = m.collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&lotto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responseLotto{
		Data: lotto,
	})
}
