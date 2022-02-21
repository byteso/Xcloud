package repository

import (
	"context"
	"fmt"

	"github.com/byteso/Xcloud/internal/database"
	"github.com/byteso/Xcloud/internal/entity"
)

type Invitation entity.Invitation

// insert
func (i *Invitation) Insert() error {
	coll := database.Database.Collection("invitation")

	result, err := coll.InsertOne(context.TODO(), i)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

// find one
func (i *Invitation) FindOne() (Invitation, error) {
	var result Invitation
	coll := database.Database.Collection("invitation")

	err := coll.FindOne(context.TODO(), i).Decode(&result)
	if err != nil {
		return Invitation{}, err
	}

	return result, nil
}
