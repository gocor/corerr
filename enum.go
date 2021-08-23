package corerr

// Err enums
const (
	ErrHTTPBadRequest     = 10400
	ErrHTTPUnauthorized   = 10401
	ErrHTTPForbidden      = 10403
	ErrHTTPNotFound       = 10404
	ErrHTTPRequestTimeout = 10408
	ErrHTTPConflict       = 10409
	ErrHTTPInternalServer = 10500
	ErrUnknown            = 99999
)

// ErrMsg Strings
const (
	ErrMsgHTTPBadRequest          = "Bad Request"
	ErrMsgHTTPUnauthorized        = "Unauthorized"
	ErrMsgHTTPForbidden           = "Forbidden"
	ErrMsgHTTPNotFound            = "Not Found"
	ErrMsgHTTPRequestTimeout      = "Request Timeout"
	ErrMsgHTTPConflict            = "Conflict"
	ErrMsgHTTPInternalServerError = "Internal Server Error"
	ErrMsgUnknown                 = "Unknown Error"
)

var errMsgMap = map[int]string{
	ErrHTTPBadRequest:     ErrMsgHTTPBadRequest,
	ErrHTTPUnauthorized:   ErrMsgHTTPUnauthorized,
	ErrHTTPForbidden:      ErrMsgHTTPForbidden,
	ErrHTTPNotFound:       ErrMsgHTTPNotFound,
	ErrHTTPRequestTimeout: ErrMsgHTTPRequestTimeout,
	ErrHTTPConflict:       ErrMsgHTTPConflict,
	ErrHTTPInternalServer: ErrMsgHTTPInternalServerError,
}

// IsErrCodeHTTP ...
func IsErrCodeHTTP(code int) bool {
	if code >= 10000 && code < 20000 {
		return true
	}
	return false
}

// ErrCodeToHTTPStatusCode returns an http status code or 0 for invalid codes
func ErrCodeToHTTPStatusCode(code int) int {
	if IsErrCodeHTTP(code) {
		return code - 10000
	}
	return 0
}
