package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/render-examples/go-gin-web-server/controllers"
	_ "github.com/render-examples/go-gin-web-server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title REST API using GO and Gin by TR
// @version 1.1
// @description This is a API Swagger documentation included book, movie and lotto API.
// @host tr-profile-go-web-server.onrender.com
// @BasePath /
// @accept json
// @produce json
// @schemes https http
// @termsOfService The Terms of Service for the API http://swagger.io/terms/.
func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go statsWorker()
}

func index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", index)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://tr-profile-go-web-server.onrender.com", "http://localhost:3010"},
		AllowMethods:     []string{"PUT", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
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
