package models_movie

type CreateMovieReq struct {
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Adult       bool    `json:"adult" bson:"adult"`
	Genres      []int   `json:"genres" bson:"genres"`
	Rating      float32 `json:"rating" bson:"rating"`
}
