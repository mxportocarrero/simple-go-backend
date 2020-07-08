package database

import (
	"context"
	"errors"
	"simple-go-backend/internal/model"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDB interface {
	CreateUser(ctx context.Context, user *model.User) (*mongo.InsertOneResult, error)
	UpdateUser(ctx context.Context, user *model.User) error
	GetAllUsers(ctx context.Context, limit string, filterKey string, filterValue string) (*mongo.Cursor, error)
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

func (db *database) GetAllUsers(ctx context.Context, limit string, filterKey string, filterValue string) (*mongo.Cursor, error) {
	var err error
	var l int
	options := options.Find()
	if limit != "" {
		l, err = strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		options.SetLimit(int64(l))
	}

	collection := db.conn.Database(db.dbName).Collection("users")

	var cursor *mongo.Cursor
	if filterKey != "" && filterValue != "" {
		cursor, err = collection.Find(ctx, bson.M{filterKey: filterValue}, options)
	}
	cursor, err = collection.Find(ctx, bson.M{}, options)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return cursor, nil
}

func (db *database) UpdateUser(ctx context.Context, user *model.User) error {
	var err error
	collection := db.conn.Database(db.dbName).Collection("users")

	filter := bson.M{"_id": user.ID}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"phone": user.Phone,
			"email": user.Email,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
