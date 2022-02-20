package types

// request struct
type RequestLogin struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestInvitation struct {
	InvitationCode string `json:"invitationCode" binding:"required"`
}

type ResponseInvitation struct {
	Account string `json:"account" binding:"required"`
}

type RequestSign struct {
	InvitationCode string `json:"invitationCode" binding:"required"`
	Account        string `json:"account" binding:"required"`
	Password       string `json:"password" binding:"required"`

	Nick   string `json:"nick"`
	Avatar string `json:"Avatar"`
}

// reponse struct
type ResponseLogin struct {
	Token string `json:"token"`
}
