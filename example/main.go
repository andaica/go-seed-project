package main

import (
	"net/http"
	"os"

	"./home"
	log "github.com/andaica/go-seed-project/logger"
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
