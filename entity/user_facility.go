package entity

import (
	"time"
)

type UserFacility struct {
	UserFacilityID  string    `gorm:"primaryKey;column:user_facility_id"`
	UserID          string    `gorm:"column:user_id"`
	FacilityLimitID string    `gorm:"column:facility_limit_id"`
	TenorID         uint      `gorm:"column:tenor"` // Sesuai tipe primary key Tenor
	Amount          int64     `gorm:"column:amount"`
	StartDate       time.Time `gorm:"column:start_date"`

	MonthlyInstallment int64 `gorm:"column:monthly_installment"`
	TotalMargin        int64 `gorm:"column:total_margin"`
	TotalPayment       int64 `gorm:"column:total_payment"`

	User          User                 `gorm:"-:all"`
	FacilityLimit UserFacilityLimit    `gorm:"-:all"`
	Tenor         Tenor                `gorm:"-:all"`
	Details       []UserFacilityDetail `gorm:"-:all"`
}

func (UserFacility) TableName() string {
	return "user_facilities"
}

type UserFacilityDetail struct {
	DetailID          string    `gorm:"primaryKey;column:detail_id"`
	UserFacilityID    string    `gorm:"column:user_facility_id"`
	DueDate           time.Time `gorm:"column:due_date"`
	InstallmentAmount int64     `gorm:"column:installment_amount"`

	UserFacility UserFacility `gorm:"-:all"`
}

func (UserFacilityDetail) TableName() string {
	return "user_facility_details"
}
