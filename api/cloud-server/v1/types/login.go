package types

type RequestLogin struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
