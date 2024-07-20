package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	mux.HandleFunc("/snippet/view", app.snippetView)

	return app.panicRecover(app.logRequest(secureHeader(mux)))

}
