package handler

import (
	"encoding/json"
	"gameReview/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ReviewHandler struct {
	revUC model.ReviewUsecase
}

func NewReviewHandler(revUC model.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{revUC}
}

func (rh *ReviewHandler) GetReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	review, err := rh.revUC.Read(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	js, err := json.Marshal(&review)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   js,
	})
}

func (rh *ReviewHandler) PostReview(c *fiber.Ctx) error {
	var review model.Review
	if err := c.BodyParser(&review); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err := rh.revUC.Add(review.Body, review.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "OK",
	})
}

func (rh *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	newBody := c.Query("body")

	err := rh.revUC.Update(id, newBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func (rh *ReviewHandler) DeleteReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	err := rh.revUC.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
