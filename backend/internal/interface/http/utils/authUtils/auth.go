package authUtils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckToken(c *fiber.Ctx, id int) bool {
	u := c.Locals("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	uid := int(claims["user"].(float64))
	return uid == id
}
