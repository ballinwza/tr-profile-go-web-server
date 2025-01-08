package services_lotto

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/connecter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ILottoService interface {
	GetAllLottoWithFilter(*gin.Context)
	GetLottoById(c *gin.Context)
	LottoScrapperByPage(*gin.Context)
}

type LottoService struct {
	collection *mongo.Collection
}

func SetupLottoService() ILottoService {
	return &LottoService{
		collection: connecter.ConnectWithMongo("tr-profile-go-db", "lotto"),
	}
}
