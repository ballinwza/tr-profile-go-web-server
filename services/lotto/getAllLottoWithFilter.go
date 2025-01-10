package services_lotto

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/handler"
	models_lotto "github.com/render-examples/go-gin-web-server/models/lotto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetAllLottoWithFilter godoc
// @Summary Response lotto list by filter
// @Description get lotto list by LottoFilterReq model then return lotto list as json
// @Tags lotties
// @Accept json
// @Produce json
// @Param year query string false "year as string number" default(2025)
// @Param b3d query string false "3 digit of string number" default(200)
// @Param b2d query string false "2 digit of string number" default(20)
// @Param f3d query string false "3 digit of string number" default(200)
// @Param first query string false "6 digit of string number" default(730209)
// @Success      200  {object}  responseLottoList
// @Failure      400  {object}  handler.ErrorResponseModel
// @Failure      404  {object}  handler.ErrorResponseModel
// @Failure      500  {object}  handler.ErrorResponseModel
// @Router       /lotto/list [get]
func (m *LottoService) GetAllLottoWithFilter(c *gin.Context) {
	var lotties []models_lotto.Lotto
	var req models_lotto.LottoFilterReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}
	filter := bson.M{}

	if req.First != nil {
		filter["first_prize"] = req.First
	}
	if req.BackThreeDigit != nil {
		filter["back_three_digit_prize"] = bson.M{"$in": req.BackThreeDigit}
	}
	if req.BackTwoDigit != nil {
		filter["back_two_digit_prize"] = req.BackTwoDigit
	}
	if req.FrontThreeDigit != nil {
		filter["front_three_digit_prize"] = bson.M{"$in": req.FrontThreeDigit}
	}
	if req.Year != nil {
		yearFromDate, err := strconv.Atoi(*req.Year)
		if err != nil {
			c.JSON(http.StatusBadRequest, handler.ErrorResponseModel{
				Error: err.Error(),
			})
			return
		}

		filter["$expr"] = bson.M{
			"$eq": []interface{}{
				bson.M{"$year": "$date"},
				yearFromDate,
			},
		}
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$sort", Value: bson.M{"date": -1}}}, // DESC
	}

	cursor, err := m.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	if err = cursor.All(context.TODO(), &lotties); err != nil {
		c.JSON(http.StatusInternalServerError, handler.ErrorResponseModel{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responseLottoList{
		Data: lotties,
	})
}
