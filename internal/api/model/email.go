package model

type Email struct {
	ID     string `json:"id" bson:"_id"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
