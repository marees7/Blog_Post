package server

import (
	"net/http"
	"time"

	"blog-post-service/src/database/mysql"
	h "blog-post-service/src/handler"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func New(dsn, host, port, serviceName string) (*http.Server, error) {
	db, err := mysql.NewDb(dsn)
	if err != nil {
		log.WithFields(log.Fields{
			"service": serviceName,
			"err":     err,
		}).Error("failed to create db")
		return nil, err
	}

	// Create new router
	router := gin.Default()

	// Get all the routes from handler.
	_, err = h.GetRoutes(router, db)
	if err != nil {
		log.WithFields(log.Fields{
			"service": serviceName,
			"err":     err,
		}).Error("failed to get http handler")
		return nil, err
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         host + ":" + port,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	return srv, nil
}
