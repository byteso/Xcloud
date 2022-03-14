package types

type ResponseStorgeInfo struct {
	TotalCapacity float64 `json:"totalCapacity" bson:"totalCapacity"`
	UsedCapacity  float64 `json:"usedCapacity" bson:"usedCapacity"`
}
