package volta

import "errors"

var (
	// ErrNoJSONMarshaler is returned when no JSON marshaler is set
	ErrorNoJSONMarshaler = errors.New("no JSON marshaler set")
	ErrorNext            = errors.New("skip to next handler")
)
