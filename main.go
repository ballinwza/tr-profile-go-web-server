package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	controllers "github.com/render-examples/go-gin-web-server/controllers"
	_ "github.com/render-examples/go-gin-web-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go statsWorker()
}

// @title My API
// @version 1.0
// @description This is a sample server for demonstrating Swagger documentation.
// @host localhost:8080

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")

	router.GET("/", index)
	controllers.SetupController().AllRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}
