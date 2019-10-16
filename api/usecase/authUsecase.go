package usecase

import (
	"context"
	"golang-websocket/api/models"
)

type AuthUsecase interface {
	Login(ctx context.Context, username string, password string) (*models.User, string, error)
}
