package mahasiswa

import (
	"context"
	"errors"
	"fmt"
	"golang-websocket/api/mocks/mahasiswa"
	"golang-websocket/api/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	mockMahasiswaRepo := new(mahasiswa.MahasiswaRepositoryMock)

	mockMahasiswa := &models.Mahasiswa{
		Nim:   "23142008",
		Nama:  "Dadang",
		Kelas: "TIB",
	}

	mockListMahasiswa := make([]*models.Mahasiswa, 0)
	mockListMahasiswa = append(mockListMahasiswa, mockMahasiswa)

	t.Run("success", func(t *testing.T) {
		mockMahasiswaRepo.On("List", mock.Anything).Return(mockListMahasiswa, nil).Once()
		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		list, err := u.List(context.TODO())

		assert.NoError(t, err)
		assert.Len(t, list, len(mockListMahasiswa))
		mockMahasiswaRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockMahasiswaRepo.On("List", mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()
		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		list, err := u.List(context.TODO())

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockMahasiswaRepo.AssertExpectations(t)
	})
}

func TestDetail(t *testing.T) {
	mockMahasiswaRepo := new(mahasiswa.MahasiswaRepositoryMock)

	mockMahasiswa := models.Mahasiswa{
		ID:    3,
		Nim:   "23142008",
		Nama:  "Dadang",
		Kelas: "TIB",
	}
	fmt.Println(mockMahasiswa)

	t.Run("success", func(t *testing.T) {
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockMahasiswa, nil).Once()

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*5)
		mahasiswa, err := u.Detail(context.TODO(), mockMahasiswa.ID)

		assert.NoError(t, err)
		assert.NotNil(t, mahasiswa)
		mockMahasiswaRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)
		mahasiswa, err := u.Detail(context.TODO(), mockMahasiswa.ID)

		assert.Error(t, err)
		assert.Nil(t, mahasiswa)
		mockMahasiswaRepo.AssertExpectations(t)

	})
}

func TestInsert(t *testing.T) {
	mockMahasiswaRepo := new(mahasiswa.MahasiswaRepositoryMock)
	mockMahasiswa := models.Mahasiswa{
		Nim:   "23142008",
		Nama:  "Dadang",
		Kelas: "TIB",
	}

	t.Run("success", func(t *testing.T) {
		tempMockMahasiswa := mockMahasiswa
		tempMockMahasiswa.ID = 0
		mockMahasiswaRepo.On("Insert", mock.Anything, mock.Anything).Return(&mockMahasiswa, nil).Once()
		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		mahasiswa, err := u.Insert(context.TODO(), tempMockMahasiswa)
		assert.NoError(t, err)
		assert.Equal(t, mockMahasiswa.Nama, tempMockMahasiswa.Nama)
		assert.NotNil(t, &mahasiswa)
		mockMahasiswaRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockMahasiswaRepo.On("Insert", mock.Anything, mock.Anything).Return(nil, errors.New("unexpected error")).Once()

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		mahasiswa, err := u.Insert(context.TODO(), mockMahasiswa)

		assert.Error(t, err)
		assert.Nil(t, mahasiswa)
		mockMahasiswaRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockMahasiswaRepo := new(mahasiswa.MahasiswaRepositoryMock)
	// var mockMhs *models.Mahasiswa
	mockMahasiswas := models.Mahasiswa{
		Nim:   "23142008",
		Nama:  "Dadang",
		Kelas: "TIB",
	}
	mockMahasiswa := make(map[string]interface{})
	mockMahasiswa["nama"] = "Dadang"
	id := int(12)

	t.Run("success", func(t *testing.T) {

		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockMahasiswas, nil).Once()
		mockMahasiswaRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Once().Return(nil)
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockMahasiswas, nil).Once()

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		mahasiswa, err := u.Update(context.TODO(), mockMahasiswa, id)

		assert.NoError(t, err)
		assert.NotNil(t, mahasiswa)
		mockMahasiswaRepo.AssertExpectations(t)
	})

	t.Run("error-in-db", func(t *testing.T) {
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error"))

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		mahasiswa, err := u.Update(context.TODO(), mockMahasiswa, id)

		assert.Error(t, err)
		assert.Nil(t, mahasiswa)
		mockMahasiswaRepo.AssertExpectations(t)
	})

	// t.Run("mahasiswa-not-found", func(t *testing.T) {
	// 	mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, nil)
	// 	mockMahasiswaRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("unexpected error")).Once()
	// 	mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, nil)

	// 	u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

	// 	mahasiswa, err := u.Update(context.TODO(), mockMahasiswa, id)

	// 	assert.Error(t, err)
	// 	assert.Nil(t, mahasiswa)
	// 	mockMahasiswaRepo.AssertExpectations(t)
	// })
}

func TestDelete(t *testing.T) {
	mockMahasiswaRepo := new(mahasiswa.MahasiswaRepositoryMock)
	mockMahasiswas := models.Mahasiswa{
		Nim:   "23142008",
		Nama:  "Dadang",
		Kelas: "TIB",
	}
	id := 3

	t.Run("success", func(t *testing.T) {
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockMahasiswas, nil).Once()
		mockMahasiswaRepo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		err := u.Delete(context.TODO(), id)

		assert.NoError(t, err)

		mockMahasiswaRepo.AssertExpectations(t)
	})

	t.Run("mahasiswa-not-found", func(t *testing.T) {
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, nil).Once()
		mockMahasiswaRepo.On("Delete", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)

		err := u.Delete(context.TODO(), id)

		assert.Nil(t, err)
		mockMahasiswaRepo.AssertExpectations(t)
	})

	t.Run("error-in-db", func(t *testing.T) {
		mockMahasiswaRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()
		u := NewMahasiswaUsecase(mockMahasiswaRepo, time.Second*2)
		err := u.Delete(context.TODO(), id)

		assert.Error(t, err)
		mockMahasiswaRepo.AssertExpectations(t)
	})
}
