package mahasiswa

import (
	"golang-websocket/api/mocks/mahasiswa"
	"golang-websocket/api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ContextMock struct {
	JSONCalled bool
}

func (c *ContextMock) JSON(code int, obj interface{}) {
	c.JSONCalled = true
}

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
	r.GET("/v1/mahasiswa/list", mahasiswaHandler.List)

	req, err := http.NewRequest(http.MethodGet, "/v1/mahasiswa/list", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockMahasiswaUsecase.AssertExpectations(t)
}
