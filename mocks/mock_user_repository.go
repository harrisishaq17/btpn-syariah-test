package mocks

import (
	"context"
	"sistem-pembiayaan/entity"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository adalah mock dari interface UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByUserIDOrPhone(ctx context.Context, userID, phone string) (*entity.User, error) {
	args := m.Called(ctx, userID, phone)

	// Untuk menghindari panic kalau return value pertama nil
	var result *entity.User
	if user, ok := args.Get(0).(*entity.User); ok {
		result = user
	}

	return result, args.Error(1)
}
