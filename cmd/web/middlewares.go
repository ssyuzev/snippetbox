package main

import (
	"log"
	"net/http"
	"os"
)

func requestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		infoLog.Printf("%s %s %s", r.Method, r.Host, r.URL)
		next(w, r)
	}
}
