package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/auth"
	"github.com/byteso/Xcloud/internal/entity"
	"github.com/byteso/Xcloud/internal/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	databaseUser       = "user"
	databaseInvitation = "invitation"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Login(request types.RequestLogin) (response types.ResponseLogin, err error) {
	var (
		r entity.UserResult
		j types.ClientJwt
	)
	filter := bson.D{{"account", request.Account}}
	result, err := repository.FindOne(databaseUser, filter)
	if err != nil {
		return response, err
	}

	err = repository.Convert(result, &r)
	if err != nil {
		return response, err
	}

	fmt.Println(r)

	if request.Account == r.Account && request.Password == r.Password {
		j.Account = r.Account
		j.BucketName = r.BucketName
		response.Token, err = auth.CreateToken(j, "client")
		fatal(err)
		fmt.Println(response.Token)
	} else {
		err = errors.New("error")
	}

	return
}

func VerifyInvitation(request types.RequestInvitation) (response types.ResponseInvitation, err error) {
	var r entity.InvitationResult

	filter := bson.D{{Key: "invitationCode", Value: request.InvitationCode}}
	result, err := repository.FindOne(databaseInvitation, filter)
	fmt.Println(result)
	if err != nil {
		return response, err
	}

	err = repository.Convert(result, &r)
	if err != nil {
		return response, err
	}

	response.Account = r.Account
	return
}

func Sign(request types.RequestSign) (err error) {
	var (
		u entity.User
		r entity.InvitationResult
	)

	filter := bson.D{{Key: "invitationCode", Value: request.InvitationCode}}
	result, err := repository.FindOne(databaseInvitation, filter)
	if err != nil {
		return
	}

	err = repository.Convert(result, &r)
	if err != nil {
		return
	}

	if r.Account == request.Account && r.Status == 0 {
		u.Account = request.Account
		u.Password = request.Password
		u.Nick = request.Nick
		// u.Avatar = request.Avatar

		uuidGen := uuid.New()
		u.BucketName = hex.EncodeToString(uuidGen[:])

		// add user
		err = repository.Insert(databaseUser, u)
		if err != nil {
			return
		}

		// add storage
		err = CreateStorage(request.Account, u.BucketName)
		if err != nil {
			return
		}

		// change invitationCode Status
		fmt.Println(r.Id)
		id, _ := primitive.ObjectIDFromHex(r.Id)
		filter := bson.D{{Key: "_id", Value: id}}
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 1}}}}

		err = repository.Update(databaseInvitation, filter, update)
		if err != nil {
			return
		}

	} else {
		err = errors.New("error")
		return
	}

	return
}
