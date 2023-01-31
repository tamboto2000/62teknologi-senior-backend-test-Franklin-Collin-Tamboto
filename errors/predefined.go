package errors

// predefined errors
var (
	ErrInternalServer          = New(ErrCodeInternalServer, "Internal Server Error", "", nil)
	ErrUnauthorizedAccessToken = New(ErrCodeUnauthorizedAccessToken, "The access token provided is not currently able to query this endpoint.", "", nil)
	ErrTokenInvalid            = New(ErrCodeTokenInvalid, "Invalid access token or authorization header.", "", nil)
)
