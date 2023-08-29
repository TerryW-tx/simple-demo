package controller

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
	// "gorm.io/gorm"
        // "gorm.io/driver/mysql"
)

func GenerateToken(username, password string) string {
	// 生成随机盐值
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, 16)
	rand.Read(salt)

	// 结合用户名、密码和盐值生成哈希摘要
	hashString := username + password + hex.EncodeToString(salt)
	md5Hash := md5.Sum([]byte(hashString))
	md5HashString := hex.EncodeToString(md5Hash[:])

	// 生成最终的唯一 Token
	token := md5HashString + hex.EncodeToString(salt)
	fmt.Println(token)
	return token
}
