package middleware

import "net/http"

type mFunc func(http.HandlerFunc) http.HandlerFunc

type Handler struct {
	handler http.HandlerFunc
}

func (h Handler) Use(m mFunc) (newHandler Handler) {
	newHandler.handler = m(h.handler)
	return
}

func (h Handler) GetHandler() http.HandlerFunc {
	return h.handler
}

func NewHandler(handler http.HandlerFunc) Handler {
	return Handler{handler: handler}
}
