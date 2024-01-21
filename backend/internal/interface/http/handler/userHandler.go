package handler

import (
	"context"
	"gameReview/internal/domain"
	"gameReview/internal/interface/http/utils/authUtils"
	"gameReview/internal/interface/http/utils/response"
	"gameReview/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUC domain.UserUsecase
	authUC usecase.AuthUC
}

func NewUserHandler(userUC domain.UserUsecase, authUC usecase.AuthUC) UserHandler {
	return UserHandler{
		userUC: userUC,
		authUC: authUC,
	}
}

func (uh *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.Err(c, 400, err.Error())
	}

	if !authUtils.CheckToken(c, id) {
		return response.Unauthorized(c)
	}

	user, err := uh.userUC.Read(id)
	if err != nil {
		return response.Err(c, 400, err.Error())
	}
	user.Password = ""
	user.ID = 0

	return response.Data(c, 400, user)
}

func (uh *UserHandler) ChangePassword(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	if !authUtils.CheckToken(c, id) {
		return response.Unauthorized(c)
	}

	p := &struct {
		Pass string `json:"password"`
	}{}

	if err := c.BodyParser(p); err != nil {
		return response.Err(c, 400, err.Error())
	}

	u, err := uh.userUC.Read(id)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	u.Password = p.Pass
	err = uh.userUC.Update(u)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	return response.Ok(c, 400, "password changed")
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	data := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(data); err != nil {
		return response.Err(c, 400, err.Error())
	}

	token, err := uh.authUC.Login(context.TODO(), data.Email, data.Password)
	if err != nil {
		return response.Err(c, 400, err.Error())
	}

	return response.Data(c, 200, token)
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	data := &struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(data); err != nil {
		return response.Err(c, 400, err.Error())
	}

	err := uh.authUC.Register(context.TODO(), data.Name, data.Email, data.Password)
	if err != nil {
		return response.Err(c, 500, err.Error())
	}

	return response.Ok(c, 200, "registered new user")
}
