package http

import (
	"gameReview/internal/transport/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(uh *handler.UserHandler) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())

	app.Get("/monitor", monitor.New(monitor.Config{Title: "App monitor"}))

	users := app.Group("/users")
	{
		users.Get("/", uh.GetUser)
	}

	return app
}
