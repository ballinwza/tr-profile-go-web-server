package services_lotto

import struct_lotto "github.com/render-examples/go-gin-web-server/models/lotto"

type responseLottoList struct {
	Data []struct_lotto.Lotto `json:"data"`
}
type responseLotto struct {
	Data struct_lotto.Lotto `json:"data"`
}
