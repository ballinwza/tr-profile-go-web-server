package handler_lotto

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	struct_lotto "github.com/render-examples/go-gin-web-server/struct/lotto"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (m *LottoService) GetAllLottoWithFilterHandler(c *gin.Context) {
	var lotties []struct_lotto.Lotto
	var req struct_lotto.LottoFilterReq
	err := c.ShouldBind(&req)
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
			c.JSON(http.StatusBadRequest, gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: err.Error(),
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
		"data": lotties,
	})
}
