package main

import (
	"net/http"
	"simple-go-backend/internal/api"
	"simple-go-backend/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug("Starting server...")

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	address := "127.0.0.1:8000"
	server := http.Server{
		Handler: router,
		Addr:    address,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed!")
	}
}
