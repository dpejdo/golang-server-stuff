package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  log.Logger
	errorLog log.Logger
}

func main() {
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{infoLog: *infoLogger, errorLog: *errorLogger}
	srv := &http.Server{
		Addr:     ":3000",
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	app.infoLog.Print("listening on port 3000")
	err := srv.ListenAndServe()
	app.errorLog.Print(err)
}
