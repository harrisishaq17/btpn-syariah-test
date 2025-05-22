package entity

type User struct {
	UserID string `gorm:"primaryKey;column:user_id"`
	Name   string `gorm:"column:name"`
	Phone  string `gorm:"column:phone"`

	FacilityLimits []UserFacilityLimit `gorm:"-:all"`
	UserFacilities []UserFacility      `gorm:"-:all"`
}

func (User) TableName() string {
	return "users"
}
