package entity

var (
	DatabaseNameUser = "user"
)

type User struct {
	Account  string `json:"account" bson:"account"`
	Password string `json:"password" bson:"password"`

	Nick   string `json:"nick" bson:"nick"`
	Avatar string `json:"avatar" bson:"avatar"`

	BucketName string `json:"bucketName" bson:"bucketName"`
}

type UserResult struct {
	Id       string `json:"_id" bson:"_id"`
	Account  string `json:"account" bson:"account"`
	Password string `json:"password" bson:"password"`

	Nick   string `json:"nick" bson:"nick"`
	Avatar string `json:"avatar" bson:"avatar"`

	BucketName string `json:"bucketName" bson:"bucketName"`
}
