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

var (
	NotFoundError    = ErrorResponse("failed", "404 Route Not Found", []string{})
	MethodNotAllowed = ErrorResponse("failed", "405 Method Not Allowed", []string{})
)
