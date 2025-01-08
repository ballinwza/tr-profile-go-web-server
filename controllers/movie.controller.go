package controllers

import "github.com/gin-gonic/gin"

func (m *Controller) MovieRoute(router *gin.Engine) {
	router.GET("/movie/list", m.GetAllMovie)
	router.GET("/movies/:id", m.GetMovieById)
	router.POST("/create/movie", m.CreateMovie)
}
