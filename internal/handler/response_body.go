package handler

type responseBody struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage any    `json:"error_message"`
	Data         any    `json:"data"`
}
