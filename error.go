package corerr

// NewStdError ...
func NewStdError(code int) error {
	return NewCoreError(code, errMsgMap[code])
}

// NewHTTPStatusError ...
func NewHTTPStatusError(httpStatusCode int) error {
	return NewStdError(10000 + httpStatusCode)
}
