package auth

import (
	"context"
	"database/sql"
	"errors"
	"golang-websocket/api/helper/authentication"
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

	user, err := a.userRepo.CheckUser(ctx, username)
	if err == sql.ErrNoRows {
		return nil, "User not found", err
	} else if err != nil {
		return nil, "", err
	}

	if authentication.ComparePassword(user.Password, password) == false {
		return nil, "", errors.New("Password Doesn't Match")
	}

	return user, "Login Success", nil
}
