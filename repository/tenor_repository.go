package repository

import (
	"sistem-pembiayaan/entity"

	"gorm.io/gorm"
)

type TenorRepository interface {
	GetAll() ([]entity.Tenor, error)
}

type tenorRepository struct {
	db *gorm.DB
}

func NewTenorRepository(db *gorm.DB) TenorRepository {
	return &tenorRepository{db}
}

func (r *tenorRepository) GetAll() ([]entity.Tenor, error) {
	var tenors []entity.Tenor
	if err := r.db.Order("tenor_value").Find(&tenors).Error; err != nil {
		return nil, err
	}
	return tenors, nil
}
