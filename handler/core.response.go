package handler

type ErrorResponseModel struct {
	Error string `json:"error" example:"Error Message"`
}

type SuccessResponseModel struct {
	Message string `json:"message" example:"Success Message"`
}
