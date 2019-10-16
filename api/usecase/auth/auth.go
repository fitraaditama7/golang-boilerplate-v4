package auth

import (
	"context"
	"database/sql"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"golang-websocket/api/usecase"
	"time"
)

type authUsecase struct {
	userRepo       repository.UserRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(userRepo repository.UserRepository, contextTimeout time.Duration) usecase.AuthUsecase {
	return &authUsecase{
		userRepo:       userRepo,
		contextTimeout: contextTimeout,
	}
}

func (a *authUsecase) Login(c context.Context, username string, password string) (*models.User, string, error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	err := a.userRepo.CheckUser(ctx, username)
	if err == sql.ErrNoRows {
		return nil, "User not found", err
	} else if err != nil {
		return nil, "", err
	}

	user, err := a.userRepo.Login(ctx, username, password)
	if err == sql.ErrNoRows {
		return nil, "Wrong password", err
	} else if err != nil {
		return nil, "", err
	}

	return user, "Login Success", nil
}
