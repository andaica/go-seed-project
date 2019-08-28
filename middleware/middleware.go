package middleware

import (
	"net/http"

	log "github.com/andaica/go-seed-project/logger"
)

type mFunc func(http.HandlerFunc) http.HandlerFunc

type Handler struct {
	handler http.HandlerFunc
}

func (h Handler) Use(m mFunc) (newHandler Handler) {
	newHandler.handler = m(h.handler)
	return
}

func NewHandler(handler http.HandlerFunc) Handler {
	return Handler{handler: handler}
}

func (h Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			log.Error("Request method incorrect")
		}
		h.handler(w, r)
	}
}

func (h Handler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Error("Request method incorrect")
		}
		h.handler(w, r)
	}
}

func (h Handler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			log.Error("Request method incorrect")
		}
		h.handler(w, r)
	}
}

func (h Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			log.Error("Request method incorrect")
		}
		h.handler(w, r)
	}
}
