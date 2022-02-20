package entity

type User struct {
	Account  string `json:"account" bson:"account"`
	Password string `json:"password" bson:"password"`

	Nick   string `json:"nick" bson:"nick"`
	Avatar string `json:"avatar" bson:"avatar"`
}
