package service

import (
	"errors"

	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	databaseNameStorage = "storage"
)

func IncreaseCapacity(account string, capacity float64) error {
	var s entity.Storage

	filter := bson.D{{"account", account}}

	result, err := repository.FindOne(databaseNameStorage, filter)
	if err != nil {
		return err
	}

	err = repository.Convert(result, &s)
	if err != nil {
		return err
	}

	if s.Account == account {
		newTotalCapacity := s.TotalCapacity + capacity
		newUsedCapacity := newTotalCapacity - s.UsedCapacity

		update := bson.D{{"$set", bson.D{{"totalCapacity", newTotalCapacity},
			{"usedCapacity", newUsedCapacity}}}}

		err := repository.Update(databaseNameStorage, filter, update)
		if err != nil {
			return err
		}

	} else {
		return errors.New("error")
	}

	return nil
}
