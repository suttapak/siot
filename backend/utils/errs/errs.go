package errs

import (
	"net/http"
)

var (
	ErrBadRequest                    = NewError(http.StatusBadRequest)                    // RFC 7231, 6.5.1
	ErrUnauthorized                  = NewError(http.StatusUnauthorized)                  // RFC 7235, 3.1
	ErrPaymentRequired               = NewError(http.StatusPaymentRequired)               // RFC 7231, 6.5.2
	ErrForbidden                     = NewError(http.StatusForbidden)                     // RFC 7231, 6.5.3
	ErrNotFound                      = NewError(http.StatusNotFound)                      // RFC 7231, 6.5.4
	ErrMethodNotAllowed              = NewError(http.StatusMethodNotAllowed)              // RFC 7231, 6.5.5
	ErrNotAcceptable                 = NewError(http.StatusNotAcceptable)                 // RFC 7231, 6.5.6
	ErrProxyAuthRequired             = NewError(http.StatusProxyAuthRequired)             // RFC 7235, 3.2
	ErrRequestTimeout                = NewError(http.StatusRequestTimeout)                // RFC 7231, 6.5.7
	ErrConflict                      = NewError(http.StatusConflict)                      // RFC 7231, 6.5.8
	ErrGone                          = NewError(http.StatusGone)                          // RFC 7231, 6.5.9
	ErrLengthRequired                = NewError(http.StatusLengthRequired)                // RFC 7231, 6.5.10
	ErrPreconditionFailed            = NewError(http.StatusPreconditionFailed)            // RFC 7232, 4.2
	ErrRequestEntityTooLarge         = NewError(http.StatusRequestEntityTooLarge)         // RFC 7231, 6.5.11
	ErrRequestURITooLong             = NewError(http.StatusRequestURITooLong)             // RFC 7231, 6.5.12
	ErrUnsupportedMediaType          = NewError(http.StatusUnsupportedMediaType)          // RFC 7231, 6.5.13
	ErrRequestedRangeNotSatisfiable  = NewError(http.StatusRequestedRangeNotSatisfiable)  // RFC 7233, 4.4
	ErrExpectationFailed             = NewError(http.StatusExpectationFailed)             // RFC 7231, 6.5.14
	ErrTeapot                        = NewError(http.StatusTeapot)                        // RFC 7168, 2.3.3
	ErrMisdirectedRequest            = NewError(http.StatusMisdirectedRequest)            // RFC 7540, 9.1.2
	ErrUnprocessableEntity           = NewError(http.StatusUnprocessableEntity)           // RFC 4918, 11.2
	ErrLocked                        = NewError(http.StatusLocked)                        // RFC 4918, 11.3
	ErrFailedDependency              = NewError(http.StatusFailedDependency)              // RFC 4918, 11.4
	ErrTooEarly                      = NewError(http.StatusTooEarly)                      // RFC 8470, 5.2.
	ErrUpgradeRequired               = NewError(http.StatusUpgradeRequired)               // RFC 7231, 6.5.15
	ErrPreconditionRequired          = NewError(http.StatusPreconditionRequired)          // RFC 6585, 3
	ErrTooManyRequests               = NewError(http.StatusTooManyRequests)               // RFC 6585, 4
	ErrRequestHeaderFieldsTooLarge   = NewError(http.StatusRequestHeaderFieldsTooLarge)   // RFC 6585, 5
	ErrUnavailableForLegalReasons    = NewError(http.StatusUnavailableForLegalReasons)    // RFC 7725, 3
	ErrInternalServerError           = NewError(http.StatusInternalServerError)           // RFC 7231, 6.6.1
	ErrNotImplemented                = NewError(http.StatusNotImplemented)                // RFC 7231, 6.6.2
	ErrBadGateway                    = NewError(http.StatusBadGateway)                    // RFC 7231, 6.6.3
	ErrServiceUnavailable            = NewError(http.StatusServiceUnavailable)            // RFC 7231, 6.6.4
	ErrGatewayTimeout                = NewError(http.StatusGatewayTimeout)                // RFC 7231, 6.6.5
	ErrHTTPVersionNotSupported       = NewError(http.StatusHTTPVersionNotSupported)       // RFC 7231, 6.6.6
	ErrVariantAlsoNegotiates         = NewError(http.StatusVariantAlsoNegotiates)         // RFC 2295, 8.1
	ErrInsufficientStorage           = NewError(http.StatusInsufficientStorage)           // RFC 4918, 11.5
	ErrLoopDetected                  = NewError(http.StatusLoopDetected)                  // RFC 5842, 7.2
	ErrNotExtended                   = NewError(http.StatusNotExtended)                   // RFC 2774, 7
	ErrNetworkAuthenticationRequired = NewError(http.StatusNetworkAuthenticationRequired) // RFC 6585, 6
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message ...string) *Error {
	err := &Error{
		Code:    code,
		Message: StatusMessage(code),
	}
	if len(message) > 0 {
		err.Message = message[0]
	}
	return err
}
func NewValidationError(message string) error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

// limits for HTTP statuscodes
const (
	statusMessageMin = 100
	statusMessageMax = 511
)

// StatusMessage returns the correct message for the provided HTTP statuscode
func StatusMessage(status int) string {
	if status < statusMessageMin || status > statusMessageMax {
		return ""
	}
	return statusMessage[status]
}

// HTTP status codes were copied from net/http.
var statusMessage = []string{
	100: "Continue",
	101: "Switching Protocols",
	102: "Processing",
	103: "Early Hints",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "Non-Authoritative Information",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	207: "Multi-Status",
	208: "Already Reported",
	226: "IM Used",
	300: "Multiple Choices",
	301: "Moved Permanently",
	302: "Found",
	303: "See Other",
	304: "Not Modified",
	305: "Use Proxy",
	306: "Switch Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	402: "Payment Required",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Request Entity Too Large",
	414: "Request URI Too Long",
	415: "Unsupported Media Type",
	416: "Requested Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}
