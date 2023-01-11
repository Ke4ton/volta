package volta

import "errors"

var (
	// ErrNoJSONMarshaler is returned when no JSON marshaler is set
	ErrorNoJSONMarshaler = errors.New("No JSON marshaler set")
)
