package handler

import (
	"encoding/json"
	"gameReview/internal/model"
	"github.com/gofiber/fiber/v2"
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
		return c.Status(404).JSON(map[string]string{"error": err.Error()})
	}

	js, err := json.Marshal(user)
	if err != nil {
		return c.Status(404).JSON(map[string]string{"error": err.Error()})
	}

	return c.Status(200).JSONP(js)
}

func (uh *UserHandler) PostUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := uh.userUC.Add(user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return c.SendStatus(201)
}
