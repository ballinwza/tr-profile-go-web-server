package controllers

import (
	"github.com/gin-gonic/gin"
	services_lotto "github.com/render-examples/go-gin-web-server/services/lotto"
	services_movie "github.com/render-examples/go-gin-web-server/services/movie"
)

type IController interface {
	AllRoute(*gin.Engine)
}
type Controller struct {
	services_lotto.ILottoService
	services_movie.IMovieService
}

func SetupController() IController {
	return &Controller{
		services_lotto.SetupLottoService(),
		services_movie.SetupMovieService(),
	}
}

func (m *Controller) AllRoute(router *gin.Engine) {
	router.GET("/lotto", m.GetLottoByIdHandler)
	router.GET("/lotto/list", m.GetAllLottoWithFilterHandler)
	router.GET("/lotto/:id", m.GetLottoByIdHandler)

	router.GET("/movies", m.GetAllMovieHandler)
	router.GET("/movies/:id", m.GetMovieByIdHandler)
	router.POST("/create/movie", m.CreateMovieHandler)
}
