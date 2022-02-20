package types

// request user info
type RequestGet struct {
	Token string `json:"token" binding:"required"`
}

type ResquestUpdata struct {
	Nick   string      `json:"nick"`
	Avatar interface{} `json:"avatar"`
}

// response user info
type ResponseGet struct {
	Account string      `json:"account"`
	Nick    string      `json:"nick"`
	Avatar  interface{} `json:"avatar"`
}
