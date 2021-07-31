package gateway_test

import (
	"content-management-api/driver/model"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockContentDriver struct {
	mock.Mock
}

func (_m *MockContentDriver) CreateModel(name string, createdAt time.Time, fields []model.Field) (*model.ContentModel, error) {
	ret := _m.Called(name, createdAt, fields)
	return ret.Get(0).(*model.ContentModel), ret.Error(1)
}

func (_m *MockContentDriver) FindContentModelByID(id string) (*model.ContentModel, error) {
	ret := _m.Called(id)
	return ret.Get(0).(*model.ContentModel), ret.Error(1)
}

func (_m *MockContentDriver) FindContentModelBySpaceID(id string) ([]model.ContentModel, error) {
	ret := _m.Called(id)
	return ret.Get(0).([]model.ContentModel), ret.Error(1)
}

func (_m *MockContentDriver) CreateEntry(entry model.Entry) (*model.Entry, error) {
	ret := _m.Called(entry)
	return ret.Get(0).(*model.Entry), ret.Error(1)
}

func (_m *MockContentDriver) DeleteContentModelByID(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
