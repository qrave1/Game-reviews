package handler

import (
	"encoding/json"
	"gameReview/internal/domain"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
	review := struct {
		Body   string `json:"body"`
		UserID int    `json:"userID"`
	}{}
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

func (rh *ReviewHandler) GetReviews(c *fiber.Ctx) error {
	reviews, err := rh.revUC.ReadAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	js, err := json.Marshal(&reviews)
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

func (rh *ReviewHandler) GetLastThree(c *fiber.Ctx) error {
	reviews, err := rh.revUC.ReadLastThree()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	js, err := json.Marshal(&reviews)
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

func (rh *ReviewHandler) UpdateReview(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	b := struct {
		Body string `json:"body"`
	}{}
	if err := c.BodyParser(&b); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err := rh.revUC.Update(id, b.Body)
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
