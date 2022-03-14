package types

// request user info
type RequestGetInfo struct {
	Account string `json:"account"`
}

type ResquestUpdateInfo struct {
	Nick   string      `json:"nick"`
	Avatar interface{} `json:"avatar"`
}

// response user info
type ResponseGetInfo struct {
	Account string      `json:"account"`
	Nick    string      `json:"nick"`
	Avatar  interface{} `json:"avatar"`
}
