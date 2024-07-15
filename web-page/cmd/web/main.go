package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog  log.Logger
	errorLog log.Logger
}

func main() {
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB("web:pass@tcp(localhost:3306)/snippetbox?parseTime=true")
	if err != nil {
		errorLogger.Fatal(err)
	}
	defer db.Close()
	app := &application{infoLog: *infoLogger, errorLog: *errorLogger}
	srv := &http.Server{
		Addr:     ":3000",
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	app.infoLog.Print("listening on port 3000")
	err = srv.ListenAndServe()
	app.errorLog.Print(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
