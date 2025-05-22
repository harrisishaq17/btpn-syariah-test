package entity

type UserFacilityLimit struct {
	FacilityLimitID string `gorm:"primaryKey;column:facility_limit_id"`
	UserID          string `gorm:"column:user_id"`
	LimitAmount     int64  `gorm:"column:limit_amount"`

	User User `gorm:"-:all"`
}

func (UserFacilityLimit) TableName() string {
	return "user_facility_limits"
}
