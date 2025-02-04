package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"fet-service/internal/service"

	"github.com/alexedwards/flow"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	logger  *log.Logger
	service *service.Service
}

func main() {
	dsn := flag.String("dsn", "./db.sqlite", "sqlite3 DSN")
	slackWebhookURL := flag.String("slack-webhook-url", "https://hooks.slack.com/services/example", "slack webhook URL for notifications")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	db, err := sql.Open("sqlite3", *dsn)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	app := &application{
		logger:  logger,
		service: &service.Service{DB: db, SlackWebhookURL: *slackWebhookURL},
	}

	mux := flow.New()
	mux.HandleFunc("/register", app.registerUserHandler, "POST")
	logger.Print("starting server on :3000")
	err = http.ListenAndServe(":3000", mux)
	logger.Fatal(err)

}
