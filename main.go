package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/mohdaalam/005/student/endpoints"
	"github.com/mohdaalam/005/student/repository"
	"github.com/mohdaalam/005/student/service"
	"github.com/mohdaalam/005/student/transport"
	"github.com/sirupsen/logrus"
)

func main() {
	dsn := "dbname='go_lang' host='localhost' user='postgres' password='root' sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)            // Set the log level to Info
	log.SetOutput(os.Stdout)                  // Set log output to standard output (stdout)
	log.SetFormatter(&logrus.TextFormatter{}) // Set the log formatter

	if err != nil {
		log.Error("unable to connect to database")
	}

	ctx := context.Background()

	var srv service.Service
	{
		repo := repository.NewRespostiry(*db, *log)
		srv = service.NewService(repo, *log)
	}
	endpoints := endpoints.NewEnpoints(srv)

	errs := make(chan error)

	fmt.Println("listening on port", 8080)
	handler := transport.NewHTTPServer(ctx, endpoints)
	errs <- http.ListenAndServe(":8080", handler)

}
