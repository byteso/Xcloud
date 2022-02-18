package types

type RequestInvitation struct {
	Account        string `json:"account" binding:"required"`
	InvitationCode string `json:"invitationCode" binding:"required"`
}
