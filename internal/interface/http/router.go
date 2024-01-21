package http

import (
	"gameReview/internal/config"
	"gameReview/internal/interface/http/handler"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(cfg *config.Config, uh handler.UserHandler, rh handler.ReviewHandler) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/monitor", monitor.New(monitor.Config{Title: "App monitor"}))

	auth := app.Group("/auth")
	{
		auth.Post("/register", uh.Register)
		auth.Post("/login", uh.Login)
	}

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(cfg.Secret),
		},
	}))

	users := app.Group("/users")
	{
		users.Get("/:id", uh.GetUser)
		users.Put("/:id", uh.ChangePassword)
	}

	reviews := app.Group("/reviews")
	{
		reviews.Get("/", rh.GetReviews)
		reviews.Get("/last", rh.GetLastThree)
		reviews.Get("/:id", rh.GetReview)
		reviews.Post("/", rh.PostReview)
		reviews.Patch("/:id", rh.UpdateReview)
		reviews.Delete("/:id", rh.DeleteReview)
	}

	return app
}
