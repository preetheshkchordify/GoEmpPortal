package models

// errorResponse defines the structure for API error responses.
// swagger:model errorResponse
type errorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Details string `json:"details,omitempty"`
}
