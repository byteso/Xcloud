package repository

import (
	"context"
	"fmt"

	"github.com/byteso/Xcloud/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

func Convert(data interface{}, value interface{}) error {
	bytes, err := bson.Marshal(data)
	if err != nil {
		return err
	}

	err = bson.Unmarshal(bytes, value)
	if err != nil {
		return err
	}

	return nil

}

func Insert(databaseName string, document interface{}) error {
	coll := database.Database.Collection(databaseName)
	_, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		return err
	}

	return nil
}

func FindOne(databaseName string, filter interface{}) (bson.D, error) {
	var result bson.D
	coll := database.Database.Collection(databaseName)
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func Find(databaseName string, filter interface{}) ([]bson.D, error) {
	var result []bson.D
	coll := database.Database.Collection(databaseName)
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return result, err
	}

	if err := cursor.All(context.TODO(), &result); err != nil {
		return result, err
	}

	fmt.Println(cursor)

	return result, err
}

func Update(databaseName string, filter, update interface{}) error {
	coll := database.Database.Collection(databaseName)

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
