package service

import (
	"errors"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/config"
)

func Login(r types.RequestLogin) (types.ResponseLogin, error) {
	if r.LoginCode == config.Config.CloudServer.LoginCode {
		// run jwt
	} else {
		return types.ResponseLogin{}, errors.New("error loginCode")
	}
	return types.ResponseLogin{}, nil
}
