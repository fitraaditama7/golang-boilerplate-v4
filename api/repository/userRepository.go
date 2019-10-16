package repository

import (
	"context"
	"golang-websocket/api/models"
)

type UserRepository interface {
	List(ctx context.Context) ([]*models.User, error)
	Detail(ctx context.Context, id int) (*models.User, error)
	Insert(ctx context.Context, user models.User) (*models.User, error)
	Update(ctx context.Context, datas map[string]interface{}, id int) error
	Delete(ctx context.Context, id int) error
	Login(ctx context.Context, username string, password string) (*models.User, error)
	CheckUser(ctx context.Context, username string) error
}
