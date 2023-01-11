package volta

type Status int

var (
	StatusSwitchingProtocols Status = 101
	StatusProcessing         Status = 102
	StatusEarlyHints         Status = 103

	StatusOK                   Status = 200
	StatusCreated              Status = 201
	StatusAccepted             Status = 202
	StatusNonAuthoritativeInfo Status = 203
	StatusNoContent            Status = 204
	StatusResetContent         Status = 205
	StatusPartialContent       Status = 206
	StatusMultiStatus          Status = 207
	StatusAlreadyReported      Status = 208

	StatusMultipleChoices   Status = 300
	StatusMovedPermanently  Status = 301
	StatusFound             Status = 302
	StatusSeeOther          Status = 303
	StatusNotModified       Status = 304
	StatusUseProxy          Status = 305
	StatusTemporaryRedirect Status = 307
	StatusPermanentRedirect Status = 308

	StatusBadRequest                  Status = 400
	StatusUnauthorized                Status = 401
	StatusPaymentRequired             Status = 402
	StatusForbidden                   Status = 403
	StatusNotFound                    Status = 404
	StatusMethodNotAllowed            Status = 405
	StatusNotAcceptable               Status = 406
	StatusProxyAuthRequired           Status = 407
	StatusRequestTimeout              Status = 408
	StatusConflict                    Status = 409
	StatusGone                        Status = 410
	StatusLengthRequired              Status = 411
	StatusPreconditionFailed          Status = 412
	StatusPayloadTooLarge             Status = 413
	StatusURITooLong                  Status = 414
	StatusUnsupportedMediaType        Status = 415
	StatusRangeNotSatisfiable         Status = 416
	StatusExpectationFailed           Status = 417
	StatusImATeapot                   Status = 418
	StatusMisdirectedRequest          Status = 421
	StatusUnprocessableEntity         Status = 422
	StatusLocked                      Status = 423
	StatusFailedDependency            Status = 424
	StatusTooEarly                    Status = 425
	StatusUpsgradeRequired            Status = 426
	StatusPreconditionRequired        Status = 428
	StatusTooManyRequests             Status = 429
	StatusRequestHeaderFieldsTooLarge Status = 431
	StatusnavailableForLegalReasons   Status = 451

	StatusInternalServerError           Status = 500
	StatusNotImplemented                Status = 501
	StatusBadGateway                    Status = 502
	StatusServiceUnavailable            Status = 503
	StatusGatewayTimeout                Status = 504
	StatusHTTPVersionNotSupported       Status = 505
	StatusVariantAlsoNegotiates         Status = 506
	StatusInsufficientStorage           Status = 507
	StatusLoopDetected                  Status = 508
	StatusNotExtended                   Status = 510
	StatusNetworkAuthenticationRequired Status = 511
)
