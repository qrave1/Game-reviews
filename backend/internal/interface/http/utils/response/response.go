package response

import "github.com/gofiber/fiber/v2"

func Ok(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).SendString(msg)
}

func Data(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status": "OK",
		"data":   data,
	})
}

func Err(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "Error",
		"message": msg,
	})
}

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(401).JSON(fiber.Map{
		"status": "Unauthorized",
	})
}
