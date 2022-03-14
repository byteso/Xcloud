package service

import (
	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func polymerizeUserInfo(userResult []entity.UserResult, storageResult []entity.Storage, result *[]types.ResponseUserInfo) {
	for _, u := range userResult {
		for _, s := range storageResult {
			if u.Account == s.Account {
				r := types.ResponseUserInfo{
					Account: u.Account,
					Nick:    u.Nick,
					Avatar:  u.Avatar,

					TotalCapacity: s.TotalCapacity,
					UsedCapacity:  s.UsedCapacity,
				}

				*result = append(*result, r)
			}
		}
	}

	return
}

func GetUserInfo(request types.RequestUserInfo) ([]types.ResponseUserInfo, error) {
	var (
		response []types.ResponseUserInfo

		userResults    []entity.UserResult
		storageResults []entity.Storage
	)

	if request.Accounts != nil {
		for _, v := range request.Accounts {
			var (
				userResult    entity.UserResult
				storageResult entity.Storage
			)
			filter := bson.D{{"account", v}}

			result, err := repository.FindOne(entity.DatabaseNameUser, filter)
			if err != nil {
				return response, err
			}

			err = repository.Convert(result, &userResult)
			if err != nil {
				return response, err
			}

			userResults = append(userResults, userResult)

			result, err = repository.FindOne(entity.DatabaseNameStorage, filter)
			if err != nil {
				return response, err
			}

			err = repository.Convert(result, &storageResult)
			if err != nil {
				return response, err
			}

			storageResults = append(storageResults, storageResult)

		}

		go polymerizeUserInfo(userResults, storageResults, &response)
	} else {
		result, err := repository.Find(entity.DatabaseNameUser, nil)
		if err != nil {
			return response, err
		}

		err = repository.Convert(result, &userResults)
		if err != nil {
			return response, err
		}

		result, err = repository.Find(entity.DatabaseNameStorage, nil)
		if err != nil {
			return response, err
		}

		err = repository.Convert(result, &storageResults)
		if err != nil {
			return response, err
		}

		go polymerizeUserInfo(userResults, storageResults, &response)
	}

	return response, nil
}
