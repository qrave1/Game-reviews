package model

type Review struct {
	ID     int    `json:"ID"`
	Body   string `json:"body"`
	UserID int    `json:"userID"`
}
