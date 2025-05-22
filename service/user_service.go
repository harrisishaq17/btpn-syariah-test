package service

import (
	"context"
	"errors"
	"net/http"
	"sistem-pembiayaan/dto/user"
	"sistem-pembiayaan/entity"
	"sistem-pembiayaan/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (int, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (int, error) {
	// Cek apakah user_id atau phone sudah ada
	if existing, _ := s.repo.FindByUserIDOrPhone(ctx, req.UserID, req.Phone); existing != nil {
		return http.StatusBadRequest, errors.New("User ID atau nomor telepon sudah digunakan.")
	}

	// Simpan user baru
	userEntity := &entity.User{
		UserID: req.UserID,
		Name:   req.Name,
		Phone:  req.Phone,
	}

	if err := s.repo.Create(ctx, userEntity); err != nil {
		return http.StatusInternalServerError, errors.New("Internal server error.")
	}

	return http.StatusCreated, nil
}
