package models_book

import "go.mongodb.org/mongo-driver/v2/bson"

type BookFilterResponse struct {
	Id            bson.ObjectID `json:"_id" bson:"_id"`
	Title         string        `json:"title" bson:"title"`
	Author        string        `json:"author" bson:"author"`
	Categories    []string      `json:"categories" bson:"categories"`
	IsBorrow      bool          `json:"is_borrow" bson:"is_borrow"`
	BorrowedCount int32         `json:"borrowed_count" bson:"borrowed_count"`
}
type BookDefaultResponse struct {
	Id            bson.ObjectID `json:"_id" bson:"_id"`
	Title         string        `json:"title" bson:"title"`
	Author        string        `json:"author" bson:"author"`
	Categories    []int32       `json:"categories" bson:"categories"`
	IsBorrow      bool          `json:"is_borrow" bson:"is_borrow"`
	BorrowedCount int32         `json:"borrowed_count" bson:"borrowed_count"`
}
