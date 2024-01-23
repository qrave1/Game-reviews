package handler

import (
	"encoding/json"
	"gameReview/internal/domain"
	"gameReview/internal/interface/http/utils/authUtils"
	"gameReview/internal/interface/http/utils/response"
	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	revUC domain.ReviewUsecase
}

func NewReviewHandler(revUC domain.ReviewUsecase) ReviewHandler {
	return ReviewHandler{revUC}
}

func (rh *ReviewHandler) GetReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	review, err := rh.revUC.Read(id)
	if err != nil {
		return response.Err(c, 400, err.Error())
	}

	if !authUtils.CheckToken(c, review.UserID) {
		return response.Unauthorized(c)
	}

	js, err := json.Marshal(&review)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	return response.Data(c, 200, js)
}

func (rh *ReviewHandler) PostReview(c *fiber.Ctx) error {
	r := struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{}
	if err := c.BodyParser(&r); err != nil {
		return response.Err(c, 400, err.Error())
	}

	uid := authUtils.GetUID(c)

	err := rh.revUC.Add(r.Title, r.Body, uid)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	return response.Ok(c, 200, "OK")
}

func (rh *ReviewHandler) GetReviews(c *fiber.Ctx) error {
	reviews, err := rh.revUC.ReadAll()
	if err != nil {
		return response.Err(c, 400, err.Error())
	}

	return response.Data(c, 200, reviews)
}

func (rh *ReviewHandler) GetLastThree(c *fiber.Ctx) error {
	reviews, err := rh.revUC.ReadLastThree()
	if err != nil {
		return response.Err(c, 400, err.Error())
	}

	return response.Data(c, 200, reviews)
}

func (rh *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	review, err := rh.revUC.Read(id)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	if !authUtils.CheckToken(c, review.UserID) {
		return response.Unauthorized(c)
	}

	b := struct {
		Body string `json:"body"`
	}{}
	if err := c.BodyParser(&b); err != nil {
		return response.Err(c, 400, err.Error())
	}

	err = rh.revUC.Update(id, b.Body)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	return response.Ok(c, 400, "updated")
}

func (rh *ReviewHandler) DeleteReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	review, err := rh.revUC.Read(id)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	if !authUtils.CheckToken(c, review.UserID) {
		return response.Unauthorized(c)
	}

	err = rh.revUC.Delete(id)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	return response.Ok(c, 200, "deleted")
}
