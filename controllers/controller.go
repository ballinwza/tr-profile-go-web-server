package controllers

import (
	"github.com/gin-gonic/gin"
	services_book "github.com/render-examples/go-gin-web-server/services/book"
	services_lotto "github.com/render-examples/go-gin-web-server/services/lotto"
	services_movie "github.com/render-examples/go-gin-web-server/services/movie"
)

type IController interface {
	AllRoute(*gin.Engine)
	LottoRoute(*gin.Engine)
	MovieRoute(*gin.Engine)
	BookRoute(*gin.Engine)
}
type Controller struct {
	services_lotto.ILottoService
	services_movie.IMovieService
	services_book.IBookService
}

func SetupController() IController {
	return &Controller{
		services_lotto.SetupLottoService(),
		services_movie.SetupMovieService(),
		services_book.SetupBookService(),
	}
}

func (m *Controller) AllRoute(router *gin.Engine) {
	m.LottoRoute(router)
	m.MovieRoute(router)
	m.BookRoute(router)
}
