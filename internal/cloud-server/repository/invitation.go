package repository

import (
	"context"
	"fmt"

	"github.com/byteso/Xcloud/internal/cloud-server/entity"
	"github.com/byteso/Xcloud/internal/database"
)

type Invitation entity.Invitation

// insert
func (i *Invitation) Insert() error {
	fmt.Println(i.Account)
	fmt.Println(i.InvitationCode)
	fmt.Println(database.Client)
	coll := database.Client.Database("XcloudTest").Collection("invitation")

	result, err := coll.InsertOne(context.TODO(), i)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
