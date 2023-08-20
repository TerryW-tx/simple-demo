package model

import (
  	// "fmt"
  	"time"
  	// "gorm.io/gorm"
)

type UserInfo struct {
  	UserId		int64
 	UserName	string
  	Password	string
	// CreateAt	time.Time `gorm:"type:DATETIME"`
	CreateAt	int64
  	Token		string
	// TokenUpdateAt	time.Time `gorm:"type:DATETIME"`
  	TokenUpdateAt	int64
  	Avatar		string
  	BackgroundImage	string
}

func CreateUserInfo(u *UserInfo) {
	db.Create(u)
}

func UpdateUserInfo(u *UserInfo) {
	var user UserInfo
	db.First(&user, "user_name = ?", (*u).UserName)
	user.Token = (*u).Token
	// user.TokenUpdateAt = u.TokenUpdateAt
	user.TokenUpdateAt = time.Now().UnixNano()
	db.Save(&user)
}

func QueryUserInfoByUsername(un string) (UserInfo, error) {
	// Query by username
	var user UserInfo
	result := db.First(&user, "user_name = ?", un)
	return user, result.Error
}

func QueryUserInfoByToken(token string) (UserInfo, error) {
	// Query by token
	var user UserInfo
	result := db.First(&user, "token = ?", token)
	return user, result.Error
}
