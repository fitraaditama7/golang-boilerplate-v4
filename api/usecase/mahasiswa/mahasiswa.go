package mahasiswa

import (
	"context"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"golang-websocket/api/usecase"
	"time"
)

type mahasiswaUsecase struct {
	mahasiswaRepo  repository.MahasiswaRepository
	contextTimeout time.Duration
}

func NewMahasiswaUsecase(mahasiswaRepo repository.MahasiswaRepository, timeout time.Duration) usecase.MahasiswaUsecase {
	return &mahasiswaUsecase{
		mahasiswaRepo:  mahasiswaRepo,
		contextTimeout: timeout,
	}
}

func (u *mahasiswaUsecase) List(c context.Context) ([]*models.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	mahasiswa, err := u.mahasiswaRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return mahasiswa, nil
}

func (u *mahasiswaUsecase) Detail(c context.Context, id int) (*models.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	mahasiswa, err := u.mahasiswaRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	return mahasiswa, nil
}

func (u *mahasiswaUsecase) Insert(c context.Context, mahasiswa models.Mahasiswa) (*models.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	mahasiswaResult, err := u.mahasiswaRepo.Insert(ctx, mahasiswa)
	if err != nil {
		return nil, err
	}
	return mahasiswaResult, nil
}

func (u *mahasiswaUsecase) Update(c context.Context, datas map[string]interface{}, id int) (*models.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.mahasiswaRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	err = u.mahasiswaRepo.Update(ctx, datas, id)
	if err != nil {
		return nil, err
	}

	mahasiswas, err := u.mahasiswaRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	return mahasiswas, nil
}

func (u *mahasiswaUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.mahasiswaRepo.Detail(ctx, id)
	if err != nil {
		return err
	}

	err = u.mahasiswaRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
