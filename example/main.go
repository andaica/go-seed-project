package main

import (
	"net/http"
	"os"

	log "../logger"
	"./home"
)

var (
	ServiceAddr = os.Getenv("GO_SERVICE_ADDR")
)

func main() {
	mux := http.NewServeMux()
	home.RegistRouter(mux)

	log.Log("server starting...")

	err := http.ListenAndServe(ServiceAddr, mux)
	if err != nil {
		log.Error("server failed to start: ", err)
	}
}
