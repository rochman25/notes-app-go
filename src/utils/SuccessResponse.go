package utils

type ResponseSuccess struct {
	Status  string
	Message string
	Data    interface{}
}

func SuccessResponse(status string, message string, data interface{}) *ResponseSuccess {
	if status == "" {
		status = "Success"
	}

	if message == "" {
		message = "Request Success"
	}

	return &ResponseSuccess{
		status,
		message,
		data,
	}
}
