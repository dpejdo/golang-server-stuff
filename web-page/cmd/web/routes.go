package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	mux.HandleFunc("/snippet/view", app.snippetView)
	standard := alice.New(app.panicRecover, app.logRequest, secureHeader)
	return standard.Then(mux)

}
