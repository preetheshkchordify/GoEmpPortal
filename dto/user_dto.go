package dto

// RegisterUserRequest represents the request body for user registration.
type RegisterUserRequest struct {
	Name        string `json:"name" example:"John Doe"`
	Email       string `json:"email" example:"john@example.com"`
	EmployeeID  string `json:"employeeId" example:"EMP123"`
	Designation string `json:"designation" example:"Software Engineer"`
	Role        string `json:"role" example:"60d21b4667d0d8992e610c85"`
	Location    string `json:"location" example:"New York"`
}

// RegisterUserResponse represents the successful response after user registration.
type RegisterUserResponse struct {
	Message  string `json:"message" example:"User registered successfully"`
	Username string `json:"username" example:"john@example.com"`
	Password string `json:"password" example:"randomGeneratedPass"`
}

// ErrorResponse represents a common error response structure.
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid input"`
}
