package config

import (
	"fmt"
	"log"
	"sistem-pembiayaan/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/spf13/viper"
)

func InitDB(v *viper.Viper) (*gorm.DB, error) {
	host := v.GetString("DB_HOST")
	port := v.GetString("DB_PORT")
	user := v.GetString("DB_USER")
	password := v.GetString("DB_PASSWORD")
	dbname := v.GetString("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	models := []interface{}{
		&entity.User{},
		&entity.UserFacilityLimit{},
		&entity.Tenor{},
		&entity.UserFacilityDetail{},
		&entity.UserFacility{},
	}

	for _, model := range models {
		err = db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Failed migrate %T: %v", model, err)
		}
		fmt.Printf("Migrated %T\n", model)
	}

	// Tambahkan foreign key constraint manual
	fkStatements := []string{
		`ALTER TABLE user_facility_limits
			ADD CONSTRAINT fk_user_facility_limits_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE`,

		`ALTER TABLE user_facilities
			ADD CONSTRAINT fk_user_facilities_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE`,

		`ALTER TABLE user_facilities
			ADD CONSTRAINT fk_user_facilities_limit FOREIGN KEY (facility_limit_id) REFERENCES user_facility_limits(facility_limit_id) ON DELETE CASCADE ON UPDATE CASCADE`,

		`ALTER TABLE user_facilities
			ADD CONSTRAINT fk_user_facilities_tenor FOREIGN KEY (tenor) REFERENCES tenors(tenor_id) ON DELETE CASCADE ON UPDATE CASCADE`,

		`ALTER TABLE user_facility_details
			ADD CONSTRAINT fk_user_facility_details_user_facility FOREIGN KEY (user_facility_id) REFERENCES user_facilities(user_facility_id) ON DELETE CASCADE ON UPDATE CASCADE`,
	}

	for _, stmt := range fkStatements {
		err := db.Exec(stmt).Error
		if err != nil {
			log.Fatalf("Gagal buat FK: %v\nStatement: %s", err, stmt)
		}
		fmt.Println("FK constraint berhasil dibuat:", stmt)
	}

	// Insert data tenor
	var count int64
	db.Model(&entity.Tenor{}).Count(&count)
	if count == 0 {
		var defaultTenors []entity.Tenor
		for i := 6; i <= 36; i += 6 {
			defaultTenors = append(defaultTenors, entity.Tenor{
				TenorValue: i,
			})
		}
		if err := db.Create(&defaultTenors).Error; err != nil {
			log.Printf("Failed to insert default tenor data: %v", err)
		}
	}

	return db, nil
}
