package dtos

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}
type SuccessResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(message string, statusCode int, detail string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Code:    statusCode,
		Details: detail,
	}
}
