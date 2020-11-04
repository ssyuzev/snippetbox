package main

import (
	"log"
	"net/http"
	"os"
)

// requestMiddleware - middleware for printing minimal request info to stdout.
func requestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		infoLog.Printf("%s %s %s", r.Method, r.Host, r.URL)
		next(w, r)
	}
}

// myMiddleware - a simple middleware example
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Any code here will execute on the way down the chain.
		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
	})
}

// secureHeaders - add some secure headers to response body.
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

// logRequest - information logger about request data.
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

/*
// demoAuthMiddleware - demo middleware for authentication.
func demoAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user isn't authorized send a 403 Forbidden status and
		// return to stop executing the chain.
		if !isAuthorised(r) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		// Otherwise, call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
*/
