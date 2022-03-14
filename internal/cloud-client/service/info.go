package service

import (
	"fmt"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func GetInfo(account string) (types.ResponseGetInfo, error) {
	var (
		r        entity.UserResult
		response types.ResponseGetInfo
	)
	filter := bson.D{{"account", account}}

	result, err := repository.FindOne(entity.DatabaseNameUser, filter)
	if err != nil {
		return response, err
	}

	err = repository.Convert(result, &r)
	if err != nil {
		return response, err
	}

	response.Account = r.Account
	response.Nick = r.Nick
	response.Avatar = r.Avatar

	return response, nil
}

func UpdateInfo(account string, request types.ResquestUpdateInfo) error {
	var (
		r entity.UserResult
	)
	filter := bson.D{{"account", account}}

	result, err := repository.FindOne(entity.DatabaseNameUser, filter)
	if err != nil {
		return err
	}

	err = repository.Convert(result, &r)
	if err != nil {
		return err
	}

	fmt.Println(r.Account)

	update := bson.D{{"$set", bson.D{{"nick", request.Nick}, {"avatar", request.Avatar}}}}
	err = repository.Update(entity.DatabaseNameUser, filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
