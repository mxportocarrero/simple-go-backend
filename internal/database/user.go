package database

import (
	"context"
	"errors"
	"simple-go-backend/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserDB interface {
	CreateUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error)
}

func (db *database) CreateUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error) {
	rating := model.Rating{
		Count:   1,
		Average: 5.0,
	}

	user.Rating = rating

	collection := db.conn.Database(db.dbName).Collection("users")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}
