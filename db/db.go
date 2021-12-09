package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=user password=password dbname=subject port=5433 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&SubjectModel{})
	return err
}

func Seed(db *gorm.DB) {
	ExecuteSeeding(db)
}
