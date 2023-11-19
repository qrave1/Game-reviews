package repository

import (
	"gameReview/internal/model"
	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (rr *ReviewRepo) Create(body string, score byte, uId int) error {
	r := model.Review{
		Body:   body,
		UserID: uId,
		Score:  score,
	}
	return rr.db.Create(&r).Error
}

func (rr *ReviewRepo) Read(id int) (model.Review, error) {
	var r model.Review
	return r, rr.db.First(&r, "id = ?", id).Error
}

func (rr *ReviewRepo) ReadAll() ([]model.Review, error) {
	var reviews []model.Review
	return reviews, rr.db.Find(&reviews).Error
}

func (rr *ReviewRepo) Update(id int, body string) error {
	return rr.db.Model(&model.Review{}).Where("id = ?", id).Update("body", body).Error
}

func (rr *ReviewRepo) Delete(id int) error {
	return rr.db.Delete(&model.Review{}, id).Error
}
