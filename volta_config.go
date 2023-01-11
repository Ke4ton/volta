package volta

// Config Volta Wrapper configuration struct
type Config struct {
	Port string
}

var (
	DefaultConfig = Config{
		Port: "8080",
	}
)
