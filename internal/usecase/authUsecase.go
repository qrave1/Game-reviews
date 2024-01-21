package usecase

import (
	"context"
	"errors"
	"fmt"
	"gameReview/internal/config"
	"gameReview/internal/domain"
	"gameReview/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	password "github.com/vzglad-smerti/password_hash"
	"time"
)

type AuthUC interface {
	Login(ctx context.Context, email, pass string) (string, error)
	Register(ctx context.Context, name, email, pass string) error
}

type AuthUsecase struct {
	cfg *config.Config
	log *logrus.Logger
	ur  domain.UserRepo
	tr  *repository.TokenRepository
}

func NewAuthUsecase(
	cfg *config.Config,
	log *logrus.Logger,
	ur domain.UserRepo,
	tr *repository.TokenRepository,
) *AuthUsecase {
	return &AuthUsecase{cfg: cfg, log: log, ur: ur, tr: tr}
}

func (au AuthUsecase) Login(ctx context.Context, email, pass string) (string, error) {
	au.log.Info("start processing login")
	defer au.log.Info("finish processing login")

	userDTO, err := au.ur.ReadByEmail(email)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	if _, err = password.Verify(userDTO.Password, pass); err != nil {
		return "", fmt.Errorf("wrong passwords: %s", pass)
	}

	token, err := au.tr.Token(ctx, email)
	if errors.Is(err, repository.ErrNotFound) {
		token, err = newJwtToken(userDTO.ID, au.cfg.Secret)
		if err != nil {
			return "", fmt.Errorf("error creating token: %w", err)
		}

		token, err = au.tr.Create(ctx, email, token)
	} else {
		if err != nil {
			return "", err
		}
	}

	return token, nil
}

func (au AuthUsecase) Register(ctx context.Context, name, email, pass string) error {
	au.log.Info("start processing register")
	defer au.log.Info("finish processing register")

	_, err := au.ur.ReadByEmail(email)
	if err == nil {
		return fmt.Errorf("user with this email already created")
	}

	hashPass, err := password.Hash(pass)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	if err := au.ur.Create(name, email, hashPass); err != nil {
		return err
	}

	u, err := au.ur.ReadByEmail(email)
	if err != nil {
		return err
	}

	token, err := newJwtToken(u.ID, au.cfg.Secret)
	if err != nil {
		return err
	}

	_, err = au.tr.Create(ctx, email, token)
	return err
}

func newJwtToken(uid int, secret string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": uid,
			"exp":  time.Now().Add(24 * time.Hour).Unix(),
		},
	)

	return token.SignedString([]byte(secret))
}
