package models

type ApiErrorResponse struct {
	Message string `json:"message"`
}

func FromString(message string) *ApiErrorResponse {
	return &ApiErrorResponse{Message: message}
}
