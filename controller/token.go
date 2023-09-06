package controller

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GenerateToken(username, password string) string {
	rand.Seed(time.Now().UnixNano())
	salt := make([]byte, 16)
	rand.Read(salt)

	hashString := username + password + hex.EncodeToString(salt)
	md5Hash := md5.Sum([]byte(hashString))
	md5HashString := hex.EncodeToString(md5Hash[:])

	token := md5HashString + hex.EncodeToString(salt)
	return token
}
