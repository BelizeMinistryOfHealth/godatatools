package models

type AccessToken struct {
	Token   string `bson:"_id" json:"token"`
	UserID  string `bson:"userId" json:"userId"`
	Deleted bool   `bson:"deleted" json:"deleted"`
}
