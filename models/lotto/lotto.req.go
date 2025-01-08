package models_lotto

type LottoFilterReq struct {
	Year            *string   `form:"year"`
	BackThreeDigit  *[]string `form:"b3d"`
	BackTwoDigit    *string   `form:"b2d"`
	First           *string   `form:"first"`
	FrontThreeDigit *[]string `form:"f3d"`
}
