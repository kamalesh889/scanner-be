package database

import (
	"barcode/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() (*model.Database, error) {

	db, err := gorm.Open(postgres.Open("postgres://postgres:kamalesh@localhost:5432/barcode"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return &model.Database{DbConn: db}, nil
}

func Migrate(db *gorm.DB) error {

	err := db.AutoMigrate(
		&model.Store{},
	)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
