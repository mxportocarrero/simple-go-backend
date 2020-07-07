package database

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO build a simple wrapper for the DB
type Database interface {
	Disconnect() error
}

type database struct {
	dbName string
	conn   *mongo.Client
}

func (db *database) Close() error {
	// TODO review this function
	// return db.conn.Disconnect()
	return nil
}

func (db *database) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := db.conn.Disconnect(ctx); err != nil {
		logrus.WithError(err).Fatal("Error disconnecting database")
		return err
	}
	return nil
}
