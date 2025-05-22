package entity

type Tenor struct {
	TenorID    uint `gorm:"primaryKey;autoIncrement;column:tenor_id"`
	TenorValue int  `gorm:"column:tenor_value"`

	UserFacilities []UserFacility `gorm:"-:all"`
}

func (Tenor) TableName() string {
	return "tenors"
}
