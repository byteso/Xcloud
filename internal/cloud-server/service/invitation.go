package service

import (
	"time"

	"github.com/byteso/Xcloud/api/cloud-server/v1/types"
	"github.com/byteso/Xcloud/internal/cloud-server/repository"
)

func AddInvitation(r types.RequestInvitation) error {
	var i repository.Invitation

	i.Account = r.Account
	i.InvitationCode = r.InvitationCode
	i.Status = 0
	i.CreatedTime = time.Now().Unix()

	if err := i.Insert(); err != nil {
		return err
	}
	return nil
}
