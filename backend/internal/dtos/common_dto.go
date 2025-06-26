package dtos

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
	Details    string `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(message string, statusCode int, detail string) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
		Details:    detail,
	}
}
