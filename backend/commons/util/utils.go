package util

type ApiResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func ApiResp(status int, message string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
