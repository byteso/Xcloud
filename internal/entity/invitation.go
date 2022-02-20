package entity

type Invitation struct {
	InvitationCode string `json:"invitationCode" bson:"invitationCode"`
	Account        string `json:"account" bson:"account"`
	Status         uint   `json:"status" bson:"status"`
	CreatedTime    int64  `json:"createdTime" bson:"createdTime"`
}
