package httpstatusdesc

import (
	"net/http"
)

const (
	StatusContinue           string = "100 Continue"
	StatusSwitchingProtocols string = "101 Switching Protocols"
	StatusProcessing         string = "102 Processing"
	StatusEarlyHints         string = "103 Early Hints"

	StatusOK                   string = "200 OK"
	StatusCreated              string = "201 Created"
	StatusAccepted             string = "202 Accepted"
	StatusNonAuthoritativeInfo string = "203 Non-Authoritative Information"
	StatusNoContent            string = "204 No Content"
	StatusResetContent         string = "205 Reset Content"
	StatusPartialContent       string = "206 Partial Content"
	StatusMultiStatus          string = "207 Multi-Status"
	StatusAlreadyReported      string = "208 Already Reported"
	StatusIMUsed               string = "226 IM Used"

	StatusMultipleChoices  string = "300 Multiple Choices"
	StatusMovedPermanently string = "301 Moved Permanently"
	StatusFound            string = "302 Found"
	StatusSeeOther         string = "303 See Other"
	StatusNotModified      string = "304 Not Modified"
	StatusUseProxy         string = "305 Use Proxy"

	StatusTemporaryRedirect string = "307 Temporary Redirect"
	StatusPermanentRedirect string = "308 Permanent Redirect"

	StatusBadRequest                   string = "400 Bad Request"
	StatusUnauthorized                 string = "401 Unauthorized"
	StatusPaymentRequired              string = "402 Payment Required"
	StatusForbidden                    string = "403 Forbidden"
	StatusNotFound                     string = "404 Not Found"
	StatusMethodNotAllowed             string = "405 Method Not Allowed"
	StatusNotAcceptable                string = "406 Not Acceptable"
	StatusProxyAuthRequired            string = "407 Proxy Authentication Required"
	StatusRequestTimeout               string = "408 Request Timeout"
	StatusConflict                     string = "409 Conflict"
	StatusGone                         string = "410 Gone"
	StatusLengthRequired               string = "411 Length Required"
	StatusPreconditionFailed           string = "412 Precondition Failed"
	StatusRequestEntityTooLarge        string = "413 Payload Too Large"
	StatusRequestURITooLong            string = "414 URI Too Long"
	StatusUnsupportedMediaType         string = "415 Unsupported Media Type"
	StatusRequestedRangeNotSatisfiable string = "416 Range Not Satisfiable"
	StatusExpectationFailed            string = "417 Expectation Failed"
	StatusTeapot                       string = "418 I'm a teapot"
	StatusMisdirectedRequest           string = "421 Misdirected Request"
	StatusUnprocessableEntity          string = "422 Unprocessable Entity"
	StatusLocked                       string = "423 Locked"
	StatusFailedDependency             string = "424 Failed Dependency"
	StatusTooEarly                     string = "425 Too Early"
	StatusUpgradeRequired              string = "426 Upgrade Required"
	StatusPreconditionRequired         string = "428 Precondition Required"
	StatusTooManyRequests              string = "429 Too Many Requests"
	StatusRequestHeaderFieldsTooLarge  string = "431 Request Header Fields Too Large"
	StatusUnavailableForLegalReasons   string = "451 Unavailable For Legal Reasons"

	StatusInternalServerError           string = "500 Internal Server Error"
	StatusNotImplemented                string = "501 Not Implemented"
	StatusBadGateway                    string = "502 Bad Gateway"
	StatusServiceUnavailable            string = "503 Service Unavailable"
	StatusGatewayTimeout                string = "504 Gateway Timeout"
	StatusHTTPVersionNotSupported       string = "505 HTTP Version Not Supported"
	StatusVariantAlsoNegotiates         string = "506 Variant Also Negotiates"
	StatusInsufficientStorage           string = "507 Insufficient Storage"
	StatusLoopDetected                  string = "508 Loop Detected"
	StatusNotExtended                   string = "510 Not Extended"
	StatusNetworkAuthenticationRequired string = "511 Network Authentication Required"
)

func Desc(code int) string {
	desc := ""

	switch code {
	case http.StatusContinue:
		desc = StatusContinue
	case http.StatusSwitchingProtocols:
		desc = StatusSwitchingProtocols
	case http.StatusProcessing:
		desc = StatusProcessing
	case http.StatusEarlyHints:
		desc = StatusEarlyHints
	case http.StatusOK:
		desc = StatusOK
	case http.StatusCreated:
		desc = StatusCreated
	case http.StatusAccepted:
		desc = StatusAccepted
	case http.StatusNonAuthoritativeInfo:
		desc = StatusNonAuthoritativeInfo
	case http.StatusNoContent:
		desc = StatusNoContent
	case http.StatusResetContent:
		desc = StatusResetContent
	case http.StatusPartialContent:
		desc = StatusPartialContent
	case http.StatusMultiStatus:
		desc = StatusMultiStatus
	case http.StatusAlreadyReported:
		desc = StatusAlreadyReported
	case http.StatusIMUsed:
		desc = StatusIMUsed
	case http.StatusMultipleChoices:
		desc = StatusMultipleChoices
	case http.StatusMovedPermanently:
		desc = StatusMovedPermanently
	case http.StatusFound:
		desc = StatusFound
	case http.StatusSeeOther:
		desc = StatusSeeOther
	case http.StatusNotModified:
		desc = StatusNotModified
	case http.StatusUseProxy:
		desc = StatusUseProxy
	case http.StatusTemporaryRedirect:
		desc = StatusTemporaryRedirect
	case http.StatusPermanentRedirect:
		desc = StatusPermanentRedirect
	case http.StatusBadRequest:
		desc = StatusBadRequest
	case http.StatusUnauthorized:
		desc = StatusUnauthorized
	case http.StatusPaymentRequired:
		desc = StatusPaymentRequired
	case http.StatusForbidden:
		desc = StatusForbidden
	case http.StatusNotFound:
		desc = StatusNotFound
	case http.StatusMethodNotAllowed:
		desc = StatusMethodNotAllowed
	case http.StatusNotAcceptable:
		desc = StatusNotAcceptable
	case http.StatusProxyAuthRequired:
		desc = StatusProxyAuthRequired
	case http.StatusRequestTimeout:
		desc = StatusRequestTimeout
	case http.StatusConflict:
		desc = StatusConflict
	case http.StatusGone:
		desc = StatusGone
	case http.StatusLengthRequired:
		desc = StatusLengthRequired
	case http.StatusPreconditionFailed:
		desc = StatusPreconditionFailed
	case http.StatusRequestEntityTooLarge:
		desc = StatusRequestEntityTooLarge
	case http.StatusRequestURITooLong:
		desc = StatusRequestURITooLong
	case http.StatusUnsupportedMediaType:
		desc = StatusUnsupportedMediaType
	case http.StatusRequestedRangeNotSatisfiable:
		desc = StatusRequestedRangeNotSatisfiable
	case http.StatusExpectationFailed:
		desc = StatusExpectationFailed
	case http.StatusTeapot:
		desc = StatusTeapot
	case http.StatusMisdirectedRequest:
		desc = StatusMisdirectedRequest
	case http.StatusUnprocessableEntity:
		desc = StatusUnprocessableEntity
	case http.StatusLocked:
		desc = StatusLocked
	case http.StatusFailedDependency:
		desc = StatusFailedDependency
	case http.StatusTooEarly:
		desc = StatusTooEarly
	case http.StatusUpgradeRequired:
		desc = StatusUpgradeRequired
	case http.StatusPreconditionRequired:
		desc = StatusPreconditionRequired
	case http.StatusTooManyRequests:
		desc = StatusTooManyRequests
	case http.StatusRequestHeaderFieldsTooLarge:
		desc = StatusRequestHeaderFieldsTooLarge
	case http.StatusUnavailableForLegalReasons:
		desc = StatusUnavailableForLegalReasons
	case http.StatusInternalServerError:
		desc = StatusInternalServerError
	case http.StatusNotImplemented:
		desc = StatusNotImplemented
	case http.StatusBadGateway:
		desc = StatusBadGateway
	case http.StatusServiceUnavailable:
		desc = StatusServiceUnavailable
	case http.StatusGatewayTimeout:
		desc = StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		desc = StatusHTTPVersionNotSupported
	case http.StatusVariantAlsoNegotiates:
		desc = StatusVariantAlsoNegotiates
	case http.StatusInsufficientStorage:
		desc = StatusInsufficientStorage
	case http.StatusLoopDetected:
		desc = StatusLoopDetected
	case http.StatusNotExtended:
		desc = StatusNotExtended
	case http.StatusNetworkAuthenticationRequired:
		desc = StatusNetworkAuthenticationRequired
	}

	return desc
}
