package corerr

// NewStdError ...
func NewStdError(code int) error {
	return NewCoreError(code, errMsgMap[code])
}

// NewWrappedStdError ...
func NewWrappedStdError(code int, err error) error {
	return NewWrappedCoreError(code, errMsgMap[code], err)
}

// NewHTTPStatusError ...
func NewHTTPStatusError(httpStatusCode int) error {
	return NewStdError(10000 + httpStatusCode)
}

// NewWrappedHTTPStatusError ...
func NewWrappedHTTPStatusError(httpStatusCode int, err error) error {
	return NewWrappedStdError(10000+httpStatusCode, err)
}
