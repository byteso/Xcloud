package login

// request struct
type RequestLogin struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestSign struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`

	Nick   string `json:"nick"`
	Avatar string `json:"Avatar"`
}

// reponse struct
type ResponseLogin struct {
	Token string `json:"token"`
}
