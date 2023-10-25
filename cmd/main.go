package main

import (
	"gameReview/internal/config"
	"gameReview/internal/database"
	"gameReview/internal/logger"
	"gameReview/internal/repository"
	"gameReview/internal/transport/http"
	"gameReview/internal/transport/http/handler"
	"gameReview/internal/usecase"
)

func main() {
	cfg := config.MustLoad()
	log := logger.NewLogger(cfg.Env)
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Error(err.Error())
	}
	ur := repository.NewUserRepo(db)
	uc := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uc)

	app := http.NewRouter(uh)

	log.Error(app.Listen(cfg.Address).Error())
}
