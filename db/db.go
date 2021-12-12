package db

import (
	"github.com/satttto/tb-micro-subject/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
	RetrieveSubjectList() ([]*model.SubjectModel, error)
}

type db struct {
	client *gorm.DB
}

func SetUpDB() (DB, error) {
	client, err := connectDB()
	if err != nil {
		return nil, err
	}

	if err := migrate(client); err != nil {
		return nil, err
	}

	return &db{
		client: client,
	}, nil
}

func connectDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=user password=password dbname=subject port=5433 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.SubjectModel{})
	return err
}

func Seed(db *gorm.DB) {
	ExecuteSeeding(db)
}

func (d *db) RetrieveSubjectList() ([]*model.SubjectModel, error) {
	subjectList := []*model.SubjectModel{}

	if err := d.client.Find(&subjectList).Error; err != nil {
		return nil, err
	}

	return subjectList, nil
}
