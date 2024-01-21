package domain

import "time"

type Review struct {
	ID        int       `json:"ID"`
	Body      string    `json:"body"`
	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}

type ReviewRepo interface {
	Create(body string, uId int) error
	Read(id int) (Review, error)
	ReadAll() ([]Review, error)
	ReadLastThree() ([]Review, error)
	Update(id int, body string) error
	Delete(id int) error
}

type ReviewUsecase interface {
	Add(body string, uId int) error
	Read(id int) (Review, error)
	ReadAll() ([]Review, error)
	ReadLastThree() ([]Review, error)
	Update(id int, body string) error
	Delete(id int) error
}
