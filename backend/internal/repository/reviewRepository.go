package repository

import (
	"gameReview/internal/domain"
	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (rr *ReviewRepo) Create(body string, uId int) error {
	r := domain.Review{
		Body:   body,
		UserID: uId,
	}

	return rr.db.Create(&r).Error
}

func (rr *ReviewRepo) Read(id int) (domain.Review, error) {
	var r domain.Review

	return r, rr.db.First(&r, "id = ?", id).Error
}

func (rr *ReviewRepo) ReadAll() ([]domain.Review, error) {
	var r []domain.Review

	return r, rr.db.Find(&r).Error
}

func (rr *ReviewRepo) ReadLastThree() ([]domain.Review, error) {
	var r []domain.Review

	return r, rr.db.Limit(3).Order("created_at desc").Find(&r).Error
}

func (rr *ReviewRepo) Update(id int, body string) error {
	return rr.db.Model(&domain.Review{}).Where("id = ?", id).Update("body", body).Error
}

func (rr *ReviewRepo) Delete(id int) error {
	return rr.db.Delete(&domain.Review{}, id).Error
}
