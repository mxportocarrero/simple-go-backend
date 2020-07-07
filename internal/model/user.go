package model

import (
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Phone     string             `json:"phone" bson:"phone"`
	Email     string             `json:"email" bson:"email"`
	Rating    Rating             `json:"rating" bson:"rating"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at"`
	Deleted   bool               `json:"-" bson:"deleted"`
}

type Rating struct {
	Count   int     `json:"count" bson:"count"`
	Average float64 `json:"average" bson:"average"`
}

func (u *User) ValidateEmail(email string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		return errors.New("Invalid Email")
	}
	return nil
}

func (u *User) ValidatePhone(phone string) error {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !re.MatchString(phone) {
		return errors.New("Invalid Phone")
	}
	return nil
}
