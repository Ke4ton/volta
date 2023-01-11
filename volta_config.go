package volta

import "encoding/json"

// Config Volta Wrapper configuration struct
type Config struct {
	Port            string
	JsonMarshaler   JSONMarshal
	JsonUnmarshaler JSONUnmarshal
}

var (
	DefaultConfig = Config{
		Port:            "8080",
		JsonMarshaler:   json.Marshal,
		JsonUnmarshaler: json.Unmarshal,
	}
)
