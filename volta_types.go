package volta

type JSONMarshal func(v interface{}) ([]byte, error)
type JSONUnmarshal func(data []byte, v interface{}) error

type Map map[string]interface{}
