package handler

import (
	"encoding/json"
	"gameReview/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserHandler struct {
	userUC model.UserUsecase
}

func NewUserHandler(userUC model.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	user, err := uh.userUC.Read(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	js, err := json.Marshal(&user)
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

func (uh *UserHandler) PostUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err := uh.userUC.Add(user.Name, user.Email, user.Password)
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

func (uh *UserHandler) ChangePassword(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	newPass := c.Query("pass")

	err := uh.userUC.UpdatePass(id, newPass)
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
