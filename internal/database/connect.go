package database

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors" // custom

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbName string
var dbLink string

func Connect() (*mongo.Client, error) {
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal("Error loading .env file")
	}
	dbName = os.Getenv("DB_NAME")
	dbLink = os.Getenv("DB_LINK")
	//dbTimeout := os.Getenv("DB_TIMEOUT")

	// connect to the database
	logrus.Debug("Connecting to the database..")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbLink))
	if err != nil {
		return nil, errors.Wrap(err, "Could not connect to de database")
	}

	// check if database is running
	// timeout, _ := strconv.Atoi(dbTimeout)
	// if err != nil {
	// 	return errors.New("error on getting database timeout")
	// }

	// if err := waitForDB(client, timeout); err != nil {
	// 	return nil, err
	// }

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // not used
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

// creates a new Database
func New() (Database, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}

	d := &database{
		dbName: dbName,
		conn:   conn,
	}

	return d, nil
}

// // TODO check this function
// func waitForDB(conn *mongo.Client, timeout int) error {
// 	ready := make(chan struct{})

// 	go func() {
// 		for {
// 			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 			defer cancel() // not used
// 			if err := conn.Ping(ctx, readpref.Primary()); err != nil {
// 				close(ready)
// 				return
// 			}
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()

// 	select {
// 	case <-ready:
// 		return nil
// 	case <-time.After(time.Duration(timeout) * time.Millisecond):
// 		return errors.New("Database not ready")
// 	}
// }
