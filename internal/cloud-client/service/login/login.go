package login

import (
	"fmt"
	"log"

	"github.com/byteso/Xcloud/api/cloud-client/v1/types"
	"github.com/byteso/Xcloud/internal/auth"
	"github.com/byteso/Xcloud/internal/repository"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Login(request types.RequestLogin) (response types.ResponseLogin, err error) {
	response.Token, err = auth.CreateToken(request, "client")
	fatal(err)
	fmt.Println(response.Token)
	return
}

func VerifyInvitation(request types.RequestInvitation) (response types.ResponseInvitation, err error) {
	var i repository.Invitation

	i.InvitationCode = request.InvitationCode

	result, err := i.FindOne()
	if err != nil {
		return response, err
	}
	response.Account = result.Account
	return
}

func Sign(request types.RequestSign) (err error) {
	return
}
