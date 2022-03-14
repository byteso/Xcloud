package types

type RequestUserInfo struct {
	Accounts []string `json:"accounts" binding:"required"`
}

type ResponseUserInfo struct {
	Account string      `json:"account"`
	Nick    string      `json:"nick"`
	Avatar  interface{} `json:"avatar"`

	TotalCapacity float64 `json:"totalCapacity" bson:"totalCapacity"`
	UsedCapacity  float64 `json:"usedCapacity" bson:"usedCapacity"`
}
