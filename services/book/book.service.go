package services_book

import (
	"github.com/gin-gonic/gin"
	"github.com/render-examples/go-gin-web-server/connecter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IBookService interface {
	GetAllBookWithFilter(*gin.Context)
	CreateBook(*gin.Context)
	UpdateBook(*gin.Context)
	BorrowBook(*gin.Context)
	GetMostBorrowBook(*gin.Context)
	DeleteBook(*gin.Context)
}

type BookService struct {
	collection *mongo.Collection
}

func SetupBookService() IBookService {
	return &BookService{
		collection: connecter.ConnectWithMongo("tr-profile-go-db", "book"),
	}
}
