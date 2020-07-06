package main

import (
	"fmt"
	"net/http"
	"os"
	"simple-go-backend/internal/api"
	"simple-go-backend/internal/config"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug(fmt.Sprintf("Starting server at localhost:%v", port))

	router, err := api.NewRouter()
	if err != nil {
		logrus.WithError(err).Fatal("Error building router")
	}

	address := fmt.Sprintf("127.0.0.1:%v", port)
	server := http.Server{
		Handler: router,
		Addr:    address,
	}

	logrus.Debug("listening..")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server failed!")
	}
}
