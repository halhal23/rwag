package db

import (
	"v1/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func Init() {
	dsn := "host=0.0.0.0 dbname=sample user=user password=password sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{})
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	d, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err = d.Close(); err != nil {
		panic(err)
	}
}