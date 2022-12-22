package utils

type Response struct {
	Status  string
	Message string
	Error   []string
}

func ErrorResponse(status string, message string, error []string) *Response {
	return &Response{
		status,
		message,
		error,
	}
}
