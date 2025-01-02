package struct_movie

type MovieRes struct {
	Uuid        string  `json:"id" bson:"id"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Adult       bool    `json:"adult" bson:"adult"`
	Genres      []int   `json:"genres" bson:"genres"`
	Rating      float32 `json:"rating" bson:"rating"`
}
