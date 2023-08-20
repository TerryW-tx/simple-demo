package model

import (
  	// "fmt"
  	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
)

var db *gorm.DB

func GetDatabase() *gorm.DB{
	// dsn := "test:123456@tcp(127.0.0.1:3306)/douyin"
  	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

func InitDatabase() error {
	var err error
	dsn := "test:123456@tcp(127.0.0.1:3306)/douyin"
  	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&VideoInfo{})

	return err
}
