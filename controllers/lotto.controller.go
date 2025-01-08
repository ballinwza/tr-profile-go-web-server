package controllers

import "github.com/gin-gonic/gin"

func (m *Controller) LottoRoute(router *gin.Engine) {
	router.GET("/lotto", m.GetLottoById)
	router.GET("/lotto/list", m.GetAllLottoWithFilter)
	router.GET("/lotto/:id", m.GetLottoById)
}
