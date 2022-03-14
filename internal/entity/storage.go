package entity

var (
	DatabaseNameStorage = "storage"
)

type Storage struct {
	Account       string  `json:"account" bson:"account"`
	BucketName    string  `json:"bucketName" bson:"bucketName"`
	TotalCapacity float64 `json:"totalCapacity" bson:"totalCapacity"`
	UsedCapacity  float64 `json:"usedCapacity" bson:"usedCapacity"`
}
