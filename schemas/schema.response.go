package schemas

type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorValidations struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
