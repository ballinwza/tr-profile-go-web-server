package controllers

import "github.com/gin-gonic/gin"

func (m *Controller) LottoRoute(router *gin.Engine) {
	router.GET("/lotto", m.GetLottoByIdHandler)
	router.GET("/lotto/list", m.GetAllLottoWithFilterHandler)
	router.GET("/lotto/:id", m.GetLottoByIdHandler)
}
