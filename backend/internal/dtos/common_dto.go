package dtos

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
	Details    string `json:"details,omitempty"`
}
