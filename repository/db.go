package repository

import (
	"github.com/go-login-prac/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDB() {
	if db != nil {
		return
	}

	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := _db.AutoMigrate(&entity.User{}); err != nil {
		panic(err)
	}
	db = _db
}

func GetDB() *gorm.DB {
	return db
}
