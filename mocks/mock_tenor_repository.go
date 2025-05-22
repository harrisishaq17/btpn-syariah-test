package mocks

import (
	"sistem-pembiayaan/entity"

	"github.com/stretchr/testify/mock"
)

type MockTenorRepository struct {
	mock.Mock
}

func (m *MockTenorRepository) GetAll() ([]entity.Tenor, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Tenor), args.Error(1)
}
