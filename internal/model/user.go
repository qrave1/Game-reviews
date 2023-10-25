package model

type User struct {
	ID       int      `json:"ID"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Reviews  []Review `json:"reviews"`
}

type UserRepo interface {
	Create(name, email, password string) error
	Read(id int) (User, error)
	UpdatePass(id int, password string) error
	Delete(id int) error
}

type UserUsecase interface {
	Add(name, email, password string) error
	Read(id int) (User, error)
	UpdatePass(id int, password string) error
	Delete(id int) error
}
