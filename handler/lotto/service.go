package handler_lotto

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/connecter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ILottoService interface {
	AllLottoServiceHandler(*gin.Engine)
	GetAllLottoHandler(*gin.Context)
	GetLottoByIdHandler(c *gin.Context)
	LottoScrapperHandler(*gin.Context)
}

type LottoService struct {
	collection *mongo.Collection
}

func SetupLottoService() ILottoService {
	return &LottoService{
		collection: connecter.ConnectWithMongo("tr-profile-go-db", "lotto"),
	}
}

func (m *LottoService) AllLottoServiceHandler(router *gin.Engine) {
	router.GET("/lotto", m.LottoScrapperHandler)
	router.GET("/lotto/list", m.GetAllLottoHandler)
	router.GET("/lotto/:id", m.GetLottoByIdHandler)
}
