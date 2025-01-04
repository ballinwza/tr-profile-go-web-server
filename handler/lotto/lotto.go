package handler_lotto

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	struct_lotto "github.com/render-examples/go-gin-web-server/struct/lotto"
)

func LottoTest(c *gin.Context) {
	var visitUrl string
	pageParam := c.DefaultQuery("page", "")
	if pageParam == "" || pageParam == "1" {
		visitUrl = "https://news.sanook.com/lotto/archive"
	} else {
		visitUrl = "https://news.sanook.com/lotto/archive/page/" + pageParam
	}

	collyCollec := colly.NewCollector()

	// ระบุว่าให้เก็บข้อมูลจาก <a> tag
	// collyCollec.OnHTML("article", func(e *colly.HTMLElement) {
	// 	// fmt.Println("Link:", e.Attr("href"))
	// 	var s *goquery.Selection
	// 	newLotto := s.Find(".archive--lotto").Text()
	// 	test = append(test, newLotto)

	// })
	var lotto []struct_lotto.Lotto

	collyCollec.OnResponse(func(r *colly.Response) {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal("Error", err)
		}

		doc.Find(".archive--lotto").Each(func(i int, s *goquery.Selection) {
			var fontThreedigitPrize, backThreeDigitPrize []string

			date, _ := s.Find("time").Attr("datetime")
			firstPrize := s.Find(".archive--lotto__result-list > li:first-child .archive--lotto__result-number").Text()

			s.Find(".archive--lotto__result-list > li:nth-child(2) .archive--lotto__result-number span").Each(func(i int, s *goquery.Selection) {
				fontThreedigitPrize = append(fontThreedigitPrize, s.Text())
			})

			s.Find(".archive--lotto__result-list > li:nth-child(3) .archive--lotto__result-number span").Each(func(i int, s *goquery.Selection) {
				backThreeDigitPrize = append(backThreeDigitPrize, s.Text())
			})

			backTwoDigitPrize := s.Find(".archive--lotto__result-list > li:nth-child(4) .archive--lotto__result-number").Text()

			lotto = append(lotto, struct_lotto.Lotto{
				Id:                   i,
				Date:                 date,
				FirstPrize:           firstPrize,
				FrontThreeDigitPrize: fontThreedigitPrize,
				BackThreeDigitPrize:  backThreeDigitPrize,
				BackTwoDigitPrize:    backTwoDigitPrize,
			})

		})

	})

	// เยี่ยมชม URL เป้าหมาย
	collyCollec.Visit(visitUrl)

	c.JSON(200, lotto)
}
