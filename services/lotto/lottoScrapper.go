package services_lotto

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	models_lotto "github.com/render-examples/go-gin-web-server/models/lotto"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (m *LottoService) LottoScrapperByPage(c *gin.Context) {
	var visitUrl string
	pageParam := c.DefaultQuery("page", "")
	if pageParam == "" || pageParam == "1" {
		visitUrl = "https://news.sanook.com/lotto/archive"
	} else {
		visitUrl = "https://news.sanook.com/lotto/archive/page/" + pageParam
	}

	collyCollec := colly.NewCollector()
	var lotties []models_lotto.Lotto

	collyCollec.OnResponse(func(r *colly.Response) {

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.Error{
				Err:  err,
				Type: http.StatusBadRequest,
				Meta: map[string]string{
					"message": "failed to get data from " + visitUrl,
				},
			})
			return
		}

		doc.Find(".archive--lotto").Each(func(i int, s *goquery.Selection) {
			var fontThreedigitPrize, backThreeDigitPrize []string

			date, _ := s.Find("time").Attr("datetime")
			timeLocation, _ := time.LoadLocation("Asia/Bangkok")
			parseDate, err := time.ParseInLocation("2006-01-02", date, timeLocation)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.Error{
					Err:  err,
					Type: http.StatusBadRequest,
					Meta: map[string]string{
						"message": "failed to convert string to time.Time",
					},
				})
				return
			}

			firstPrize := s.Find(".archive--lotto__result-list > li:first-child .archive--lotto__result-number").Text()

			s.Find(".archive--lotto__result-list > li:nth-child(2) .archive--lotto__result-number span").Each(func(i int, s *goquery.Selection) {
				fontThreedigitPrize = append(fontThreedigitPrize, s.Text())
			})

			s.Find(".archive--lotto__result-list > li:nth-child(3) .archive--lotto__result-number span").Each(func(i int, s *goquery.Selection) {
				backThreeDigitPrize = append(backThreeDigitPrize, s.Text())
			})

			backTwoDigitPrize := s.Find(".archive--lotto__result-list > li:nth-child(4) .archive--lotto__result-number").Text()

			lotties = append(lotties, models_lotto.Lotto{
				Date:                 parseDate,
				FirstPrize:           firstPrize,
				FrontThreeDigitPrize: fontThreedigitPrize,
				BackThreeDigitPrize:  backThreeDigitPrize,
				BackTwoDigitPrize:    backTwoDigitPrize,
			})

		})

	})

	collyCollec.Visit(visitUrl)

	var operation []mongo.WriteModel
	opts := options.BulkWrite().SetOrdered(false)

	for _, lotto := range lotties {
		updateFilter := mongo.NewUpdateOneModel().SetFilter(bson.M{"date": lotto.Date}).SetUpdate(bson.M{"$set": lotto}).SetUpsert(true)
		operation = append(operation, updateFilter)

	}

	res, err := m.collection.BulkWrite(context.TODO(), operation, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.Error{
			Err:  err,
			Type: http.StatusInternalServerError,
			Meta: map[string]string{
				"message": "failed to insert lotto into DB",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "lotties was updated",
		"response": res,
	})
}
