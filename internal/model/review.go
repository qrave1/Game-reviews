package model

type Review struct {
	ID     int    `json:"ID"`
	Body   string `json:"body"`
	Score  byte   `json:"score"`
	UserID int    `json:"userID"`
}

type ReviewRepo interface {
	Create(body string, score byte, uId int) error
	Read(id int) (Review, error)
	ReadAll() ([]Review, error)
	Update(id int, body string) error
	Delete(id int) error
}

type ReviewUsecase interface {
	Add(body string, score byte, uId int) error
	Read(id int) (Review, error)
	ReadAll() ([]Review, error)
	Update(id int, body string) error
	Delete(id int) error
}
