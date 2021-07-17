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

func Seed() {
	users := []entity.User{
		{Email: "hoge@example.com", Password: "hogehoge"},
		{Email: "hoge2@example.com", Password: "hogehoge2"},
		{Email: "hoge3@example.com", Password: "hogehoge3"},
	}
	if err := db.Create(users).Error; err != nil {
		panic(err)
	}
}
