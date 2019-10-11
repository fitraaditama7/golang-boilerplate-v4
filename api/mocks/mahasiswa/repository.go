package mahasiswa

import (
	"context"
	"golang-websocket/api/models"

	"github.com/stretchr/testify/mock"
)

type MahasiswaRepositoryMock struct {
	mock.Mock
}

func (_m *MahasiswaRepositoryMock) List(ctx context.Context) ([]*models.Mahasiswa, error) {
	ret := _m.Called(ctx)

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

func (_m *MahasiswaRepositoryMock) Detail(ctx context.Context, id int) (*models.Mahasiswa, error) {
	ret := _m.Called(ctx)

	var r0 *models.Mahasiswa
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.Mahasiswa); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Mahasiswa)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MahasiswaRepositoryMock) Insert(ctx context.Context, mahasiswa models.Mahasiswa) (*models.Mahasiswa, error) {
	ret := _m.Called(ctx)

	var r0 *models.Mahasiswa
	if rf, ok := ret.Get(0).(func(context.Context, models.Mahasiswa) *models.Mahasiswa); ok {
		r0 = rf(ctx, mahasiswa)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Mahasiswa)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Mahasiswa) error); ok {
		r1 = rf(ctx, mahasiswa)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MahasiswaRepositoryMock) Update(ctx context.Context, datas map[string]interface{}, id int) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, int) error); ok {
		r0 = rf(ctx, datas, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MahasiswaRepositoryMock) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
