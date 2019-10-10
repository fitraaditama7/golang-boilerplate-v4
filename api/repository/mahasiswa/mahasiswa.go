package mahasiswa

import (
	"context"
	"database/sql"
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

	query := `SELECT id, nim, nama, kelas FROM mahasiswa`

	rows, err := m.Conn.QueryContext(ctx, query)
	// defer func() {
	// 	err := rows.Close()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }()
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
	query := `SELECT id, nim, nama, kelas FROM mahasiswa WHERE id = ` + ids
	err := m.Conn.QueryRowContext(ctx, query).Scan(&mahasiswa.ID, &mahasiswa.Nim, &mahasiswa.Nama, &mahasiswa.Kelas)
	if err != nil {
		return nil, err
	}
	return &mahasiswa, err
}

func (m *mysqlMahasiswaRepository) Insert(ctx context.Context, mahasiswa models.Mahasiswa) (*models.Mahasiswa, error) {
	query := `INSERT INTO mahasiswa (nim, nama, kelas) VALUES(?, ?, ?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx, mahasiswa.Nim, mahasiswa.Nama, mahasiswa.Kelas)
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

	var data []interface{}
	for index, value := range datas {
		if value != "" {
			query = query + " " + index + ` = ?, `
			data = append(data, value)
		}
	}
	query = strings.TrimRight(query, ", ")

	id := strconv.Itoa(ids)
	query = query + " WHERE id = " + id

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	_, err = stmt.ExecContext(ctx, data...)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (m *mysqlMahasiswaRepository) Delete(ctx context.Context, id int) error {
	ids := strconv.Itoa(id)
	query := `DELETE FROM mahasiswa WHERE id = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, query, ids)
	if err != nil {
		return err
	}
	return nil
}
