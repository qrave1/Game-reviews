package repository

import (
	"gameReview/internal/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Create(name, email, password string) error {
	u := domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return ur.db.Create(&u).Error
}

func (ur *UserRepo) Read(id int) (domain.User, error) {
	var u domain.User
	return u, ur.db.First(&u, "id = ?", id).Error
}

func (ur *UserRepo) ReadByEmail(email string) (domain.User, error) {
	var u domain.User
	return u, ur.db.First(&u, "email = ?", email).Error
}

func (ur *UserRepo) Update(user domain.User) error {
	return ur.db.Model(domain.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (ur *UserRepo) Delete(id int) error {
	return ur.db.Delete(&domain.User{}, id).Error
}
