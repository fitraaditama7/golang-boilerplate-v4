package mahasiswa

import (
	"context"
	"database/sql"
	"fmt"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"strconv"
	"strings"
)

type mysqlMahasiswaRepository struct {
	Conn *sql.DB
}

func NewMahasiswaRepository(Conn *sql.DB) repository.MahasiswaRepository {
	return &mysqlMahasiswaRepository{Conn}
}

func (m *mysqlMahasiswaRepository) List(ctx context.Context) ([]*models.Mahasiswa, error) {
	var mahasiswas []*models.Mahasiswa

	query := `SELECT * FROM mahasiswa`

	rows, err := m.Conn.QueryContext(ctx, query)
	defer func() {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := new(models.Mahasiswa)
		err := rows.Scan(
			&row.ID,
			&row.Nim,
			&row.Nama,
			&row.Kelas,
		)
		if err != nil {
			return nil, err
		}
		mahasiswas = append(mahasiswas, row)
	}
	return mahasiswas, nil
}

func (m *mysqlMahasiswaRepository) Detail(ctx context.Context, id int) (*models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	ids := strconv.Itoa(id)
	query := `SELECT * FROM mahasiswa WHERE id = ` + ids
	err := m.Conn.QueryRowContext(ctx, query).Scan(&mahasiswa.ID, &mahasiswa.Nim, &mahasiswa.Nama, &mahasiswa.Kelas)
	if err != nil {
		return nil, err
	}
	return &mahasiswa, err
}

func (m *mysqlMahasiswaRepository) Insert(ctx context.Context, mahasiswa models.Mahasiswa) (*models.Mahasiswa, error) {
	query := fmt.Sprintf(`INSERT INTO mahasiswa (nim, nama, kelas) VALUES("%s", "%s", "%s")`, mahasiswa.Nim, mahasiswa.Nama, mahasiswa.Kelas)
	res, err := m.Conn.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	mahasiswa.ID = int(id)
	return &mahasiswa, nil
}

func (m *mysqlMahasiswaRepository) Update(ctx context.Context, datas map[string]interface{}, ids int) (int, error) {
	query := `UPDATE mahasiswa SET`
	for index, value := range datas {
		if value != "" {
			query = query + " " + index + ` = "` + value.(string) + `", `
		}
	}
	query = strings.TrimRight(query, ", ")

	id := strconv.Itoa(ids)
	query = query + " WHERE id = " + id

	_, err := m.Conn.ExecContext(ctx, query)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (m *mysqlMahasiswaRepository) Delete(ctx context.Context, id int) error {
	ids := strconv.Itoa(id)
	query := `DELETE FROM mahasiswa WHERE id = ` + ids
	_, err := m.Conn.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
