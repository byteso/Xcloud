package service

import (
	"errors"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/auth"
	"github.com/byteso/Xcloud/internal/config"
)

func Login(r types.RequestLogin) (types.ResponseLogin, error) {
	var response types.ResponseLogin
	if r.LoginCode == config.Config.CloudServer.LoginCode {
		// run jwt
		token, err := auth.CreateToken(r, "server")
		if err != nil {
			return response, err
		}

		response.Token = token
	} else {
		return response, errors.New("error loginCode")
	}

	return response, nil
}
