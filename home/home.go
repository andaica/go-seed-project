package home

import (
	"net/http"

	log "../logger"
	"../middleware"
)

const message = "Hello seed!"

func RegistRouter(mux *http.ServeMux) {
	// mux.HandleFunc("/", log.TimeLogging(HomePage))
	// when use middleware, must type:
	homeHandler := middleware.NewHandler(HomePage)
	mux.HandleFunc("/", homeHandler.Use(log.TimeLogging).GetHandler())
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
