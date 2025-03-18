package errors

import "fmt"

// AppError represents a custom application error
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Common error types
var (
	ErrInvalidInput = &AppError{
		Code:    400,
		Message: "Invalid input data",
	}

	ErrNotFound = &AppError{
		Code:    404,
		Message: "Resource not found",
	}

	ErrUnauthorized = &AppError{
		Code:    401,
		Message: "Unauthorized access",
	}

	ErrForbidden = &AppError{
		Code:    403,
		Message: "Access forbidden",
	}

	ErrInternalServer = &AppError{
		Code:    500,
		Message: "Internal server error",
	}

	ErrDuplicateEntry = &AppError{
		Code:    409,
		Message: "Resource already exists",
	}
)

// WrapError wraps an existing error with additional context
func WrapError(err error, message string) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return &AppError{
			Code:    appErr.Code,
			Message: message,
			Err:     appErr,
		}
	}

	return &AppError{
		Code:    500,
		Message: message,
		Err:     err,
	}
}
