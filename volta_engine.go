package volta

import (
	"fmt"
	"net/http"
)

// Config Volta Wrapper configuration struct
type Config struct {
	Port string
}

var DefaultConfig = Config{
	Port: "8080",
}

type Engine struct {
	Conf Config

	http *http.Server
}

func New(conf Config) *Engine {
	if conf.Port == "" {
		conf = DefaultConfig
	}

	return &Engine{
		Conf: conf,
	}
}

func (e *Engine) SetConfig(conf Config) {
	e.Conf = conf
}

func (e *Engine) Run() {
	e.http = &http.Server{
		Addr: ":" + e.Conf.Port,
	}

	go func() {
		if err := http.ListenAndServe(e.http.Addr, nil); err != nil {
			panic(err)
		}
	}()

	fmt.Println("[Volta] Server started on port :" + e.Conf.Port)

	select {}
}
