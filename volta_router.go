package volta

import (
	"github.com/pearleascent/volta/dependencies/httprouter"
	"net/http"
)

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request

	ps httprouter.Params

	jsonMarshaler   JSONMarshal
	jsonUnmarshaler JSONUnmarshal
}

type Handler func(*Ctx) error

func (a *App) Get(path string, handler ...Handler) {
	a.router.GET(path, a.wrap(handler))
}

func (a *App) Post(path string, handler ...Handler) {
	a.router.POST(path, a.wrap(handler))
}

func (a *App) Put(path string, handler ...Handler) {
	a.router.PUT(path, a.wrap(handler))
}

func (a *App) Delete(path string, handler ...Handler) {
	a.router.DELETE(path, a.wrap(handler))
}

func (a *App) Patch(path string, handler ...Handler) {
	a.router.PATCH(path, a.wrap(handler))
}

func (a *App) Options(path string, handler ...Handler) {
	a.router.OPTIONS(path, a.wrap(handler))
}

func (a *App) wrap(handlers []Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := &Ctx{
			Response:        w,
			Request:         r,
			ps:              ps,
			jsonUnmarshaler: a.Conf.JsonUnmarshaler,
			jsonMarshaler:   a.Conf.JsonMarshaler,
		}

		for _, middlewares := range a.using {
			if err := middlewares(ctx); err != nil {
				return
			}
		}

		for _, handler := range handlers {
			if err := handler(ctx); err == nil {
				return
			} else if err != ErrorNext {

			} else {
				ctx.Response.Write([]byte(err.Error()))
			}
		}
	}
}
