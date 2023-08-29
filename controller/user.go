package controller

import (
	"fmt"
	"time"
	"context"
	// "github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/RaymondCode/simple-demo/dal"
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

var userIdSequence = int32(1)
var ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func GenerateUserId() int32 {
	atomic.AddInt32(&userIdSequence, 1)
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
	u := dal.User
	_, err := u.WithContext(ctx).Where(u.Username.Eq(username), u.Password.Eq(password)).Take()
	fmt.Println("query success")

	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
                })
	} else {
		fmt.Println("new user")
		user := entity.User{
			UserID: GenerateUserId(),
			Username: username,
			Password: password,
			CreateTime: time.Now().Unix(),
			Token: token,
			TokenUpdateTime: time.Now().Unix(),
			Avatar: "teststring",
			BackgroundImage: "testimage",
			FollowCount: 0,
			FollowerCount: 0,
			FavoriteCount: 0,
			WorkCount: 0,
		}
		// model.CreateUserInfo(&new_user)
		u.WithContext(ctx).Create(&user)
		c.JSON(
			http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 0},
                        UserId:   int64(user.UserID),
                        Token:    user.Token,
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
	userDal := dal.User
	user, err := userDal.WithContext(ctx).Where(userDal.Username.Eq(username), userDal.Password.Eq(password)).Take()
	fmt.Println("Login Success")

	if err == nil {
		// user.Token = GenerateToken(username, password)
		// user.TokenUpdateTime = time.Now().Unix()
		fmt.Println(user.Token)
		// model.UpdateUserInfo(&user)
                // userDal.WithContext(ctx).Where(userDal.Username.Eq(username), userDal.Password.Eq(password)).UpdateSimple(
		// 	userDal.Token.Value(user.Token), 
		// 	userDal.TokenUpdateTime.Value(user.TokenUpdateTime),
		// )
		c.JSON(http.StatusOK, UserLoginResponse{
                        Response: Response{StatusCode: 0},
                        UserId:   int64(user.UserID),
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

	userDal := dal.User
	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	fmt.Println("Get Userinfo Success")
	fmt.Println(user)
	fmt.Println(err)

	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:  User{
					Id:            user.UserID,
					Name:          user.Username,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
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
