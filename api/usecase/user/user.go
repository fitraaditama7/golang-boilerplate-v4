package user

import (
	"context"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"golang-websocket/api/usecase"
	"time"
)

type userUsecase struct {
	userRepo       repository.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepo repository.UserRepository, contextTimeout time.Duration) usecase.UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		contextTimeout: contextTimeout,
	}
}

func (u *userUsecase) List(c context.Context) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	users, err := u.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) Detail(c context.Context, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Insert(c context.Context, user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	users, err := u.userRepo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) Update(c context.Context, datas map[string]interface{}, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.userRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.Update(ctx, datas, id)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (u *userUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.userRepo.Detail(ctx, id)
	if err != nil {
		return err
	}

	err = u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
