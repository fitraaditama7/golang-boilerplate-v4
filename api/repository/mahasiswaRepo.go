package repository

import (
	"context"
	"golang-websocket/api/models"
)

type MahasiswaRepository interface {
	List(ctx context.Context) ([]*models.Mahasiswa, error)
	Detail(ctx context.Context, id int) (*models.Mahasiswa, error)
	Insert(ctx context.Context, mahasiswa models.Mahasiswa) (*models.Mahasiswa, error)
	Update(ctx context.Context, datas map[string]interface{}, id int) (int, error)
	Delete(ctx context.Context, id int) error
}
