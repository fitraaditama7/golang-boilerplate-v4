package mahasiswa

import (
	"context"
	"golang-websocket/api/models"

	"github.com/stretchr/testify/mock"
)

type MockMahasiswaUsecase struct {
	mock.Mock
}

func (m *MockMahasiswaUsecase) List(ctx context.Context) ([]*models.Mahasiswa, error) {
	ret := m.Called()

	var r0 []*models.Mahasiswa
	if rf, ok := ret.Get(0).(func(context.Context) []*models.Mahasiswa); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Mahasiswa)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockMahasiswaUsecase) Detail(ctx context.Context, id int) (*models.Mahasiswa, error) {
	ret := m.Called()

	var r0 *models.Mahasiswa
	if rf, ok := ret.Get(0).(func(ctx context.Context, id int) *models.Mahasiswa); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Mahasiswa)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, id int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockMahasiswaUsecase) Insert(ctx context.Context, mahasiswa models.Mahasiswa) (*models.Mahasiswa, error) {
	ret := m.Called()

	var r0 *models.Mahasiswa
	if rf, ok := ret.Get(0).(func(ctx context.Context, mahasiswa models.Mahasiswa) *models.Mahasiswa); ok {
		r0 = rf(ctx, mahasiswa)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Mahasiswa)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, mahasiswa models.Mahasiswa) error); ok {
		r1 = rf(ctx, mahasiswa)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockMahasiswaUsecase) Update(ctx context.Context, datas map[string]interface{}, id int) (*models.Mahasiswa, error) {
	ret := m.Called()

	var r0 *models.Mahasiswa
	if rf, ok := ret.Get(0).(func(ctx context.Context, datas map[string]interface{}, id int) *models.Mahasiswa); ok {
		r0 = rf(ctx, datas, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Mahasiswa)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctx context.Context, datas map[string]interface{}, id int) error); ok {
		r1 = rf(ctx, datas, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockMahasiswaUsecase) Delete(ctx context.Context, id int) error {
	ret := m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func(ctx context.Context, id int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
