package logger

import (
	"log"
	"net/http"
	"time"
)

func TimeLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer log.Printf("request processed in: %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func Error(v ...interface{}) {
	log.Println("Error: ", v)
}
