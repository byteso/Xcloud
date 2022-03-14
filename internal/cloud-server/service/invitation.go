package service

import (
	"time"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	databaseName = "invitation"
)

func CreateInvitation(r types.RequestInvitation) error {
	var i entity.Invitation

	i.Account = r.Account
	i.InvitationCode = r.InvitationCode
	i.Status = 0
	i.CreatedTime = time.Now().Unix()

	if err := repository.Insert(databaseName, i); err != nil {
		return err
	}
	return nil
}

func polymerizeInvitationInfo(userResult []entity.UserResult, invitationResult []entity.InvitationResult, result *[]types.ResponseInvitationInfo) {
	for _, u := range userResult {
		for _, i := range invitationResult {
			if u.Account == i.Account {
				r := types.ResponseInvitationInfo{
					Account: u.Account,

					InvitationCode: i.InvitationCode,
					Status:         i.Status,
				}

				*result = append(*result, r)
			}
		}
	}

	return
}

func GetInvitationInfo(request types.RequestInvitationInfo) ([]types.ResponseInvitationInfo, error) {
	var (
		response []types.ResponseInvitationInfo

		userResults       []entity.UserResult
		invitationResults []entity.InvitationResult
	)

	if request.Accounts != nil {
		for _, v := range request.Accounts {
			var (
				userResult       entity.UserResult
				InvitationResult entity.InvitationResult
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

			result, err = repository.FindOne(entity.DatabaseNameInvitation, filter)
			if err != nil {
				return response, err
			}

			err = repository.Convert(result, &InvitationResult)
			if err != nil {
				return response, err
			}

			invitationResults = append(invitationResults, InvitationResult)

		}

		go polymerizeInvitationInfo(userResults, invitationResults, &response)
	} else {
		result, err := repository.Find(entity.DatabaseNameUser, nil)
		if err != nil {
			return response, err
		}

		err = repository.Convert(result, &userResults)
		if err != nil {
			return response, err
		}

		result, err = repository.Find(entity.DatabaseNameInvitation, nil)
		if err != nil {
			return response, err
		}

		err = repository.Convert(result, &invitationResults)
		if err != nil {
			return response, err
		}

		go polymerizeInvitationInfo(userResults, invitationResults, &response)
	}

	return response, nil
}

func DeleteInvitation(r types.RequestInvitation) error {
	return nil
}
