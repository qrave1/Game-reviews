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

	rr := repository.NewReviewRepo(db)
	rc := usecase.NewReviewUsecase(rr)
	rh := handler.NewReviewHandler(rc)

	app := http.NewRouter(uh, rh)

	log.Error(app.Listen(cfg.Address))

	// todo написать unit тесты для всего
	// todo сделать фронт
	// todo сделать auth и middleware
}
