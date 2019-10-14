package mahasiswa

import (
	"encoding/json"
	"golang-websocket/api/mocks/mahasiswa"
	"golang-websocket/api/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockListMahasiswa := []*models.Mahasiswa{
		&models.Mahasiswa{
			ID: 1, Nim: "23142008", Nama: "Dadang", Kelas: "TIB",
		},
		&models.Mahasiswa{
			ID: 2, Nim: "23142009", Nama: "Dudung", Kelas: "TIB",
		},
	}
	mockMahasiswaUsecase := new(mahasiswa.MockMahasiswaUsecase)
	mockMahasiswaUsecase.On("List", mock.Anything).Return(mockListMahasiswa, nil)

	mahasiswaHandler := MahasiswaHandler{
		MahasiswaUsecase: mockMahasiswaUsecase,
	}

	r := gin.Default()
	r.GET("/list", mahasiswaHandler.List)

	req, err := http.NewRequest(http.MethodGet, "/list", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockMahasiswaUsecase.AssertExpectations(t)
}

func TestDetail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockMahasiswa := models.Mahasiswa{
		ID: 1, Nim: "23142008", Nama: "Dadang", Kelas: "TIB",
	}

	mockMahasiswaUsecase := new(mahasiswa.MockMahasiswaUsecase)
	mockMahasiswaUsecase.On("Detail", mock.Anything, mock.Anything).Return(&mockMahasiswa, nil).Once()

	mahasiswaHandler := MahasiswaHandler{
		MahasiswaUsecase: mockMahasiswaUsecase,
	}

	r := gin.Default()
	r.GET("/detail/:id", mahasiswaHandler.Detail)

	req, err := http.NewRequest(http.MethodGet, "/detail/"+strconv.Itoa(mockMahasiswa.ID), nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockMahasiswaUsecase.AssertExpectations(t)
}

func TestInsert(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockMahasiswa := models.Mahasiswa{
		ID: 1, Nim: "23142008", Nama: "Dadang", Kelas: "TIB",
	}

	tempMockMahasiswa := mockMahasiswa
	tempMockMahasiswa.ID = 0
	j, err := json.Marshal(tempMockMahasiswa)
	assert.NoError(t, err)

	mockMahasiswaUsecase := new(mahasiswa.MockMahasiswaUsecase)

	mockMahasiswaUsecase.On("Insert", mock.Anything, mock.Anything).Return(&mockMahasiswa, nil)

	mahasiswaHandler := MahasiswaHandler{
		MahasiswaUsecase: mockMahasiswaUsecase,
	}

	r := gin.Default()
	r.POST("/insert", mahasiswaHandler.Insert)

	req, err := http.NewRequest(http.MethodPost, "/insert", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockMahasiswaUsecase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockMahasiswaUsecase := new(mahasiswa.MockMahasiswaUsecase)

	mockMahasiswaUsecase.On("Delete", mock.Anything, mock.Anything).Return(nil)

	mahasiswaHandler := MahasiswaHandler{
		MahasiswaUsecase: mockMahasiswaUsecase,
	}

	r := gin.Default()
	r.DELETE("/delete/:id", mahasiswaHandler.Delete)
	id := 10
	req, err := http.NewRequest(http.MethodDelete, "/delete/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockMahasiswaUsecase.AssertExpectations(t)
}
