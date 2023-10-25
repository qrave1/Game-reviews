package usecase

import "gameReview/internal/model"

type UserUsecase struct {
	userRepo model.UserRepo
}

func NewUserUsecase(userRepo model.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uu UserUsecase) Add(name, email, password string) error {
	return uu.userRepo.Create(name, email, password)
}

func (uu UserUsecase) Read(id int) (model.User, error) {
	return uu.userRepo.Read(id)
}

func (uu UserUsecase) UpdatePass(id int, password string) error {
	return uu.userRepo.UpdatePass(id, password)
}

func (uu UserUsecase) Delete(id int) error {
	return uu.userRepo.Delete(id)
}
