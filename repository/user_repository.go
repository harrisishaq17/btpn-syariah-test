package repository

import (
    "context"
    "sistem-pembiayaan/entity"

    "gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByUserIDOrPhone(ctx context.Context, userID, phone string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByUserIDOrPhone(ctx context.Context, userID string, phone string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Where("user_id = ? OR phone = ?", userID, phone).
		First(&user).Error
        
	if err != nil {
		return nil, err
	}

	return &user, nil
}
