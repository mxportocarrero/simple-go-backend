package main

import (
	"fmt"
	"net/http"
	"os"
	"simple-go-backend/internal/api"
	"simple-go-backend/internal/config"
	"simple-go-backend/internal/database"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// ENVIROMENT VARIABLES
	// ====================
	var err error
	err = godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}

	// DATABASE CONNECTION
	// ====================
	db, err := database.New()
	if err != nil {
		logrus.WithError(err).Fatal("Error connecting Database")
	}

	defer func(db database.Database) {
		if err := db.Disconnect(); err != nil {
			logrus.WithError(err).Fatal("Error disconnecting database")
		}
	}(db)

	// SERVER SET UP
	// ====================
	port := os.Getenv("PORT")

	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("version", config.Version).Debug(fmt.Sprintf("Starting server at localhost:%v", port))

	router, err := api.NewRouter(db)
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
