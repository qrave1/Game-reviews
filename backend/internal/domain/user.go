package domain

type User struct {
	ID       int      `json:"ID,omitempty"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password,omitempty"`
	Reviews  []Review `json:"reviews,omitempty"`
}

type UserRepo interface {
	Create(name, email, password string) error
	Read(id int) (User, error)
	ReadByEmail(email string) (User, error)
	Update(user User) error
	Delete(id int) error
}

type UserUsecase interface {
	Add(name, email, password string) error
	Read(id int) (User, error)
	ReadByEmail(email string) (User, error)
	Update(user User) error
	Delete(id int) error
}
