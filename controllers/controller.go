package controllers

import (
	"github.com/gin-gonic/gin"
	services_lotto "github.com/render-examples/go-gin-web-server/services/lotto"
)

type IController interface {
	AllRoute(*gin.Engine)
}
type Controller struct {
	services_lotto.ILottoService
}

func SetupController() IController {
	return &Controller{
		services_lotto.SetupLottoService(),
	}
}

func (m *Controller) AllRoute(router *gin.Engine) {
	router.GET("/lotto", m.GetLottoByIdHandler)
	router.GET("/lotto/list", m.GetAllLottoWithFilterHandler)
	router.GET("/lotto/:id", m.GetLottoByIdHandler)
}
