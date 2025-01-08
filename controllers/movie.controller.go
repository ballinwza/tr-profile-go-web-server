package controllers

import "github.com/gin-gonic/gin"

func (m *Controller) MovieRoute(router *gin.Engine) {
	router.GET("/movies", m.GetAllMovieHandler)
	router.GET("/movies/:id", m.GetMovieByIdHandler)
	router.POST("/create/movie", m.CreateMovieHandler)
}
