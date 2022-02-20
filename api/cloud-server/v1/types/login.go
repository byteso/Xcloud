package types

type RequestLogin struct {
	LoginCode string `json:"loginCode" binding:"required"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}
