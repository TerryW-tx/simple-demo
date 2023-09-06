package controller

import (
	"context"
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// context to be developed
var ctx = context.TODO()

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// generate by random haskey and md5
	token := GenerateToken(username, password)
	userDal := dal.User
	_, err := userDal.WithContext(ctx).
		Where(userDal.Username.Eq(username), userDal.Password.Eq(password)).
		Take()

	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		user := entity.User{
			// UserID: null,
			Username:        username,
			Password:        password,
			CreateTime:      time.Now().Unix(),
			Token:           token,
			TokenUpdateTime: time.Now().Unix(),
			Avatar:          "teststring",
			BackgroundImage: "testimage",
			FollowCount:     0,
			FollowerCount:   0,
			FavoriteCount:   0,
			WorkCount:       0,
		}
		userDal.WithContext(ctx).Create(&user)
		c.JSON(
			http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.UserID,
				Token:    user.Token,
			})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userDal := dal.User
	user, err := userDal.WithContext(ctx).
		Where(userDal.Username.Eq(username), userDal.Password.Eq(password)).
		Take()

	if err == nil {
		user.Token = GenerateToken(username, password)
		user.TokenUpdateTime = time.Now().Unix()
		userDal.WithContext(ctx).
			Where(userDal.Username.Eq(username), userDal.Password.Eq(password)).
			UpdateSimple(
				userDal.Token.Value(user.Token),
				userDal.TokenUpdateTime.Value(user.TokenUpdateTime),
			)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.UserID,
			Token:    user.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userDal := dal.User
	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()

	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *ConvertUserEntityToController(user, userId),
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
