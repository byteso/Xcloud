package service

import (
	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/config"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateStorage(account string, bucketName string) error {
	var s entity.Storage

	s.Account = account
	s.BucketName = bucketName
	s.TotalCapacity = config.Config.CloudClient.DefaultStorage
	s.UsedCapacity = 0

	err := repository.Insert(entity.DatabaseNameStorage, s)
	if err != nil {
		return err
	}

	err = CreateBucket(bucketName)
	if err != nil {
		return err
	}

	return nil
}

func GetStorageInfo(account string) (types.ResponseStorgeInfo, error) {
	var (
		r        entity.Storage
		response types.ResponseStorgeInfo
	)
	filter := bson.D{{"account", account}}

	result, err := repository.FindOne(entity.DatabaseNameStorage, filter)
	if err != nil {
		return response, err
	}

	err = repository.Convert(result, &r)
	if err != nil {
		return response, err
	}

	response.TotalCapacity = r.TotalCapacity
	response.UsedCapacity = r.UsedCapacity

	return response, nil

}

func DeleteStorage(account string) error {
	return nil
}
