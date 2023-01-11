package volta

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)
type Handlers []Handler

type Routes interface {
}
