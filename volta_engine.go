package volta

import (
	"fmt"
	"github.com/pearleascent/volta/dependencies/httprouter"
	"net/http"
)

type App struct {
	Conf Config

	http   *http.Server
	router *httprouter.Router

	using []Handler
}

func New(conf Config) *App {
	if conf.Port == "" {
		conf.Port = DefaultConfig.Port
	}
	if conf.JsonUnmarshaler == nil {
		conf.JsonUnmarshaler = DefaultConfig.JsonUnmarshaler
	}
	if conf.JsonMarshaler == nil {
		conf.JsonMarshaler = DefaultConfig.JsonMarshaler
	}

	return &App{
		Conf:   conf,
		router: httprouter.New(),
	}
}

func (a *App) Run() {
	a.http = &http.Server{
		Addr: ":" + a.Conf.Port,
	}

	go func() {
		if err := http.ListenAndServe(a.http.Addr, a.router); err != nil {
			panic(err)
		}
	}()

	fmt.Println("[Volta] Server started on port :" + a.Conf.Port)

	select {}
}

func (a *App) Use(handlers ...Handler) {
	a.using = append(a.using, handlers...)
}

func (a *App) Stop() {
	if err := a.http.Close(); err != nil {
		panic(err)
	}
}
