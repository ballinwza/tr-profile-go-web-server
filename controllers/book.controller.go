package controllers

import "github.com/gin-gonic/gin"

func (c *Controller) BookRoute(router *gin.Engine) {
	router.GET("/book/list", c.GetAllBookWithFilter)
	router.POST("/create/book", c.CreateBook)
	router.PUT("/update/book/:id", c.UpdateBook)
	router.PUT("/borrow/book/:id", c.BorrowBook)
	router.GET("/book/popular", c.GetMostBorrowBook)
	router.DELETE("/delete/book/:id", c.DeleteBook)
}
