package controller

import (
	"fmt"
	"time"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func GenerateUserId() int64 {
	atomic.AddInt64(&userIdSequence, 1)
	return userIdSequence
}

// func GenerateToken(username, password string) string {
// 	return username + password
// }

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// token := username + password
	token := GenerateToken(username, password)

	if exist, _ := ValidateToken(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
                })
	} else {
		fmt.Println("new user")
		new_user := model.UserInfo{
			UserId: GenerateUserId(),
			UserName: username,
			Password: password,
			CreateAt: time.Now().UnixNano(),
			Token: token,
			TokenUpdateAt: time.Now().UnixNano(),
			Avatar: "teststring",
			BackgroundImage: "testimage",
		}
		model.CreateUserInfo(&new_user)
		c.JSON(
			http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 0},
                        UserId:   new_user.UserId,
                        Token:    new_user.Token,
		})
	}

	/*
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
	*/
}

// func TokenValidation() {
// 
// }

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// token := username + password
	// token := GenerateToken(username, password)

	if user, exist := model.QueryUserInfoByUsername(username); exist == nil && password == user.Password {
		user.Token = GenerateToken(username, password)
		user.TokenUpdateAt = time.Now().UnixNano()
		fmt.Println(user.Token)
		model.UpdateUserInfo(&user)
                c.JSON(http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 0},
                        UserId:   user.UserId,
                        Token:    user.Token,
                })
        } else {
		c.JSON(http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
                })
	}

	/*
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	*/
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if _, exist := model.QueryUserInfoByToken(token); exist == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:  User{
					Id:            1,
					Name:          "zhanglei",
					FollowCount:   10,
					FollowerCount: 5,
					IsFollow:      true,
			}, 
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	/*
	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	*/
}
