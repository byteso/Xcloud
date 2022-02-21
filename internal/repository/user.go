package repository

import (
	"context"

	"github.com/byteso/Xcloud/internal/database"
	"github.com/byteso/Xcloud/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type User entity.User

func (u *User) Insert() error {
	coll := database.Database.Collection("user")
	_, err := coll.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) FindOne() error {
	var result User
	coll := database.Database.Collection("user")
	err := coll.FindOne(context.TODO(), u).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
	}

	return nil
}

func (u *User) Find(filter interface{}) error {
	coll := database.Database.Collection("user")
	_, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update() {

}
