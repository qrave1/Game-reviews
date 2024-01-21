package handler

import (
	"context"
	"gameReview/internal/domain"
	"gameReview/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": "wrong credentials",
	})

	user, err := uh.userUC.Read(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	user.Password = ""
	user.ID = 0

	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})
}

func (uh *UserHandler) ChangePassword(c *fiber.Ctx) error {
	id := c.QueryInt("id")

	p := &struct {
		Pass string `json:"password"`
	}{}

	if err := c.BodyParser(p); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	u, err := uh.userUC.Read(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	u.Password = p.Pass
	err = uh.userUC.Update(u)
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

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	data := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	token, err := uh.authUC.Login(context.TODO(), data.Email, data.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"token":  token,
	})
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	data := &struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err := uh.authUC.Register(context.TODO(), data.Name, data.Email, data.Password)
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

//func checkToken() bool {
//	u := c.Locals("user").(*jwt.Token)
//	claims := u.Claims.(jwt.MapClaims)
//	uid := int(claims["user"].(float64))
//	if uid != id
//}
