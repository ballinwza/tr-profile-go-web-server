package struct_movie

type BindMovieById struct {
	Id string `uri:"id" binding:"required"`
}

type CreateMovieReq struct {
	Uuid        *string `json:"id,omitempty" bson:"id,omitempty"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description"`
	Adult       bool    `json:"adult" bson:"adult"`
	Genres      []int   `json:"genres" bson:"genres"`
	Rating      float32 `json:"rating" bson:"rating"`
}
