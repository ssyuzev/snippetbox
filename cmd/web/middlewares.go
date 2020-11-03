package main

import (
	"log"
	"net/http"
)

func requestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.Host, r.URL)
		next(w, r)
	}
}
