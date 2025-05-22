package service_test

import (
	"context"
	"errors"
	"sistem-pembiayaan/dto/user"
	"sistem-pembiayaan/entity"
	"sistem-pembiayaan/mocks"
	"sistem-pembiayaan/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	req := &user.CreateUserRequest{
		UserID: "U123",
		Name:   "John Doe",
		Phone:  "08123456789",
	}

	// Simulasikan tidak ditemukan user dengan userID/phone yang sama
	mockRepo.On("FindByUserIDOrPhone", mock.Anything, req.UserID, req.Phone).
		Return(nil, gorm.ErrRecordNotFound)

	// Simulasikan penyimpanan user berhasil
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(nil)

	status, err := userService.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Duplicate(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	req := &user.CreateUserRequest{
		UserID: "U123",
		Name:   "Jane Doe",
		Phone:  "08123456789",
	}

	// Simulasikan user dengan ID/phone tersebut sudah ada
	mockRepo.On("FindByUserIDOrPhone", mock.Anything, req.UserID, req.Phone).
		Return(&entity.User{UserID: "U123", Phone: "08123456789"}, nil)

	status, err := userService.CreateUser(context.Background(), req)

	assert.Equal(t, 400, status)
	assert.EqualError(t, err, "User ID atau nomor telepon sudah digunakan.")
}

func TestCreateUser_RepoErrorOnCreate(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	req := &user.CreateUserRequest{
		UserID: "U456",
		Name:   "Bob Smith",
		Phone:  "0822334455",
	}

	// Tidak ditemukan duplikat user
	mockRepo.On("FindByUserIDOrPhone", mock.Anything, req.UserID, req.Phone).
		Return(nil, gorm.ErrRecordNotFound)

	// Simulasi error saat menyimpan user
	mockRepo.On("Create", mock.Anything, mock.Anything).
		Return(errors.New("DB error"))

	status, err := userService.CreateUser(context.Background(), req)

	assert.Equal(t, 500, status)
	assert.EqualError(t, err, "Internal server error.")
}
