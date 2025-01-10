package models_book

type BookRequest struct {
	Title         string  `json:"title" bson:"title" form:"title" query:"title" binding:"required"`
	Author        string  `json:"author" bson:"author" form:"author" query:"author" binding:"required"`
	Categories    []int32 `json:"categories" bson:"categories" form:"categories" query:"categories" binding:"required"`
	IsBorrow      bool    `json:"is_borrow" bson:"is_borrow" form:"is_borrow" query:"is_borrow"`
	BorrowedCount int32   `json:"borrowed_count" bson:"borrowed_count" form:"borrowed_count" query:"borrowed_count"`
}
type BookUpdateRequest struct {
	Title      string  `json:"title" bson:"title" form:"title" query:"title"`
	Author     string  `json:"author" bson:"author" form:"author" query:"author"`
	Categories []int32 `json:"categories" bson:"categories" form:"categories" query:"categories"`
}

type BookFilterRequest struct {
	Title      *string  `form:"title" query:"title"`
	Author     *string  `form:"author" query:"author"`
	Categories *[]int32 `form:"categories" query:"categories"`
}
