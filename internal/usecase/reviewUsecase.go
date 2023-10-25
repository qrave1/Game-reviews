package usecase

import "gameReview/internal/model"

type ReviewUsecase struct {
	reviewRepo model.ReviewRepo
}

func NewReviewUsecase(reviewRepo model.ReviewRepo) *ReviewUsecase {
	return &ReviewUsecase{reviewRepo: reviewRepo}
}

func (ru *ReviewUsecase) Add(body string, uId int) error {
	return ru.reviewRepo.Create(body, uId)
}

func (ru *ReviewUsecase) Read(id int) (model.Review, error) {
	return ru.reviewRepo.Read(id)
}

func (ru *ReviewUsecase) Update(id int, body string) error {
	return ru.reviewRepo.Update(id, body)
}

func (ru *ReviewUsecase) Delete(id int) error {
	return ru.reviewRepo.Delete(id)
}
