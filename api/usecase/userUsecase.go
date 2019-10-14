package usecase

import (
	"context"
	"golang-websocket/api/models"
)

type UserUsecase interface {
	List(ctx context.Context) ([]*models.User, error)
	Detail(ctx context.Context, id int) (*models.User, error)
	Insert(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, datas map[string]interface{}, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
}
