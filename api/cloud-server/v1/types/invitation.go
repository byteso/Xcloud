package types

type RequestInvitation struct {
	Account        string `json:"account" binding:"required"`
	InvitationCode string `json:"invitationCode" binding:"required"`
}

type RequestInvitationInfo struct {
	Accounts []string `json:"account" binding:"required"`
}

type ResponseInvitationInfo struct {
	Account        string `json:"account"`
	InvitationCode string `json:"invitationCode"`
	Status         uint   `json:"status"`
}
