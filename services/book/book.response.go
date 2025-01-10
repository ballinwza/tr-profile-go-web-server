package services_book

import models_book "github.com/render-examples/go-gin-web-server/models/book"

type responseBookList struct {
	Data []models_book.BookFilterResponse `json:"data"`
}
type responseBook struct {
	Data models_book.BookFilterResponse `json:"data"`
}

type bookCategories struct {
	Action  int32 `exmaple:"1"`
	SciFi   int32 `exmaple:"2"`
	Fantasy int32 `exmaple:"3"`
	Drama   int32 `exmaple:"4"`
}
