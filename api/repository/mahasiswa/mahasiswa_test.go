package mahasiswa

import (
	"context"
	"golang-websocket/api/models"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	mockMahasiswa := []*models.Mahasiswa{
		&models.Mahasiswa{
			ID: 1, Nim: "23142008", Nama: "Dadang", Kelas: "TIB",
		},
		&models.Mahasiswa{
			ID: 2, Nim: "23142009", Nama: "Dudung", Kelas: "TIB",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "nim", "nama", "kelas"}).
		AddRow(mockMahasiswa[0].ID, mockMahasiswa[0].Nim, mockMahasiswa[0].Nama, mockMahasiswa[0].Kelas).
		AddRow(mockMahasiswa[1].ID, mockMahasiswa[1].Nim, mockMahasiswa[1].Nama, mockMahasiswa[1].Kelas)
	query := `SELECT id, nim, nama, kelas FROM mahasiswa`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewMahasiswaRepository(db)
	mahasiswa, err := a.List(context.TODO())
	assert.NoError(t, err)
	assert.Len(t, mahasiswa, 2)
}

func TestDetail(t *testing.T) {
	var num = 1
	id := strconv.Itoa(num)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was nnot exptected when opening stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "nim", "nama", "kelas"}).
		AddRow(1, "23142008", "Dadang", "TIB")

	query := "SELECT id, nim, nama, kelas FROM mahasiswa WHERE id = " + id
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewMahasiswaRepository(db)
	mahasiswa, err := a.Detail(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, mahasiswa)
}

func TestInsert(t *testing.T) {
	mahasiswa := models.Mahasiswa{
		Nim:   "23142008",
		Nama:  "Dadang",
		Kelas: "TIB",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was nnot exptected when opening stub database connection", err)
	}

	query := `INSERT INTO mahasiswa (nim, nama, kelas) VALUES(?, ?, ?)`
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(mahasiswa.Nim, mahasiswa.Nama, mahasiswa.Kelas).WillReturnResult(sqlmock.NewResult(9, 1))
	a := NewMahasiswaRepository(db)
	mahasiswaa, err := a.Insert(context.TODO(), mahasiswa)
	assert.NoError(t, err)
	assert.Equal(t, 9, mahasiswaa.ID)
}

func TestUpdate(t *testing.T) {
	var datas = make(map[string]interface{})
	datas["nama"] = "Dadangs"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was nnot exptected when opening stub database connection", err)
	}
	query := `UPDATE mahasiswa SET`
	var data []interface{}
	for index, value := range datas {
		if value != "" {
			query = query + " " + index + ` = \?, `
			data = append(data, value)
		}
	}
	query = strings.TrimRight(query, ", ")

	ids := 3
	id := strconv.Itoa(ids)
	query = query + " WHERE id = " + id

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("Dadangs").WillReturnResult(sqlmock.NewResult(3, 1))

	a := NewMahasiswaRepository(db)
	err = a.Update(context.TODO(), datas, ids)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was nnot exptected when opening stub database connection", err)
	}

	query := `DELETE FROM mahasiswa WHERE id = ?`
	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WillReturnResult(sqlmock.NewResult(12, 1))
	num := 12
	a := NewMahasiswaRepository(db)
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}
