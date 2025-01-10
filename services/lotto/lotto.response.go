package services_lotto

import models_lotto "github.com/render-examples/go-gin-web-server/models/lotto"

type responseLottoList struct {
	Data []models_lotto.Lotto `json:"data"`
}
type responseLotto struct {
	Data models_lotto.Lotto `json:"data"`
}
