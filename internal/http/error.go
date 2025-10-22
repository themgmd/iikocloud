package http

// ErrorResponse is a custom error type that represents iikoCloud API Error.
type ErrorResponse struct {
	StatusCode       int    `json:"-"`
	CorrelationID    string `json:"correlationId"`
	ErrorDescription string `json:"errorDescription"`
	ErrorField       string `json:"error"`
}

// Error ...
func (e ErrorResponse) Error() string {
	return e.ErrorDescription
}
