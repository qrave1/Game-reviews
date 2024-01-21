package usecase

import "gameReview/internal/domain"

type UserUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(userRepo domain.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uu *UserUsecase) Add(name, email, password string) error {
	return uu.userRepo.Create(name, email, password)
}

func (uu *UserUsecase) Read(id int) (domain.User, error) {
	return uu.userRepo.Read(id)
}

func (uu *UserUsecase) ReadByEmail(email string) (domain.User, error) {
	return uu.userRepo.ReadByEmail(email)
}

func (uu *UserUsecase) Update(user domain.User) error {
	return uu.userRepo.Update(user)
}

func (uu *UserUsecase) Delete(id int) error {
	return uu.userRepo.Delete(id)
}
