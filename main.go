package main

import (
	"blog-post-service/src/database/mysql"
	"blog-post-service/src/server"
	"blog-post-service/src/utils/constants"
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func start() {
	log.Info("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Warn("failed to load env")
	}

	srv, err := server.New(mysql.DSN(), os.Getenv("HOST"), os.Getenv("PORT"), constants.ServiceName)
	if err != nil {
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Fatal("failed to create http server")
	}

	log.WithFields(log.Fields{
		"service": constants.ServiceName,
	}).Info("starting server")

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.WithFields(log.Fields{
			"service": constants.ServiceName,
			"err":     err,
		}).Fatal("failed to listen on http server")
	}
}

func main() {
	start()
}
