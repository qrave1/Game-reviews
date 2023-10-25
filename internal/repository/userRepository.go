package repository

import (
	"gameReview/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur UserRepo) Create(name, email, password string) error {
	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
		Reviews:  nil,
	}
	return ur.db.Create(&user).Error
}

func (ur UserRepo) Read(id int) (model.User, error) {
	var user model.User
	return user, ur.db.First(&user, "id = ?", id).Error
}

func (ur UserRepo) UpdatePass(id int, password string) error {
	return ur.db.Model(&model.User{}).Where("id = ?", id).Update("password", password).Error

}

func (ur UserRepo) Delete(id int) error {
	return ur.db.Delete(&model.User{}, id).Error
}
