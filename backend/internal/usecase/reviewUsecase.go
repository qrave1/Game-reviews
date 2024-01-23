package usecase

import "gameReview/internal/domain"

type ReviewUsecase struct {
	reviewRepo domain.ReviewRepo
}

func NewReviewUsecase(reviewRepo domain.ReviewRepo) *ReviewUsecase {
	return &ReviewUsecase{reviewRepo: reviewRepo}
}

func (ru *ReviewUsecase) Add(title, body string, uId int) error {
	return ru.reviewRepo.Create(title, body, uId)
}

func (ru *ReviewUsecase) Read(id int) (domain.Review, error) {
	return ru.reviewRepo.Read(id)
}

func (ru *ReviewUsecase) ReadAll() ([]domain.Review, error) {
	return ru.reviewRepo.ReadAll()
}

func (ru *ReviewUsecase) ReadLastThree() ([]domain.Review, error) {
	return ru.reviewRepo.ReadLastThree()
}

func (ru *ReviewUsecase) Update(id int, body string) error {
	return ru.reviewRepo.Update(id, body)
}

func (ru *ReviewUsecase) Delete(id int) error {
	return ru.reviewRepo.Delete(id)
}
