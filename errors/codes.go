package errors

// general purpose codes
const (
	ErrCodeValidation              = "VALIDATION_ERROR"
	ErrCodeInternalServer          = "INTERNAL_SERVER_ERROR"
	ErrCodeUnauthorizedAccessToken = "UNAUTHORIZED_ACCESS_TOKEN"
	ErrCodeTokenInvalid            = "TOKEN_INVALID"
)

// specific case codes
const (
	ErrBusinessNotFound = "BUSINESS_NOT_FOUND"
)
