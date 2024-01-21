//go:build wireinject
// +build wireinject

package wire

import (
	"gameReview/internal/config"
	"gameReview/internal/domain"
	"gameReview/internal/interface/http"
	"gameReview/internal/interface/http/handler"
	"gameReview/internal/logger"
	"gameReview/internal/repository"
	"gameReview/internal/storage/mysql"
	"gameReview/internal/storage/redis"
	"gameReview/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type Container struct {
	Router *fiber.App
	cfg    *config.Config
}

func NewContainer(router *fiber.App, cfg *config.Config) *Container {
	return &Container{
		Router: router,
		cfg:    cfg,
	}
}

func (c *Container) Run() error {
	err := c.Router.Listen(c.cfg.Address)
	if err != nil {
		return err
	}
	return nil
}

func InitializeContainer() (*Container, error) {
	wire.Build(NewContainer,
		http.NewRouter,
		handler.NewUserHandler, handler.NewReviewHandler,
		wire.Bind(new(domain.UserUsecase), new(*usecase.UserUsecase)),
		wire.Bind(new(domain.ReviewUsecase), new(*usecase.ReviewUsecase)),
		wire.Bind(new(usecase.AuthUC), new(*usecase.AuthUsecase)),
		usecase.NewUserUsecase, usecase.NewReviewUsecase, usecase.NewAuthUsecase,
		wire.Bind(new(domain.UserRepo), new(*repository.UserRepo)),
		wire.Bind(new(domain.ReviewRepo), new(*repository.ReviewRepo)),
		repository.NewUserRepo, repository.NewReviewRepo, repository.NewTokenRepository,
		mysql.NewDB, redis.NewRedisStorage,
		config.NewDefault, logger.NewLogger)
	return &Container{}, nil
}
