package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestMiddleware(app.home))
	mux.HandleFunc("/snippet", requestMiddleware(app.showSnippet))
	mux.HandleFunc("/snippet/create", requestMiddleware(app.createSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
