package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessAndResponseJson(data interface{}) (int, map[string]interface{}) {
	return http.StatusOK, gin.H{"data": data}
}
