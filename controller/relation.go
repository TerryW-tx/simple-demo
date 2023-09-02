package controller

import (
	// "fmt"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/RaymondCode/simple-demo/dal"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType := c.Query("action_type")

	userDal := dal.User
	byUser, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	toUser, err := userDal.WithContext(ctx).Where(userDal.UserID.Eq(toUserId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if actionType == "1" {
		err = CreateFollow(c, byUser, toUser)
	} else if actionType == "2" {
		err = CancelFollow(c, byUser, toUser)
	}

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}


func CreateFollow(c *gin.Context, toUser, byUser *entity.User) error {
	toUserId := toUser.UserID
	byUserId := byUser.UserID
	
	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		follow := entity.Follow{
			FollowbyID: toUserId,
			FollowerID: byUserId,
			CreateTime: time.Now().Unix(),
		}
		followDal := dal.Follow
		userDal := dal.User
		err := followDal.WithContext(ctx).Create(&follow)
		if err != nil {
			return err
		}
		_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(toUserId)).UpdateSimple(userDal.FollowerCount.Add(1))
		if err != nil {
			return err
		}
		_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(byUserId)).UpdateSimple(userDal.FollowCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	})

	if err == nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Follow success"},
		})
	}

	return err
}

func CancelFollow(c *gin.Context, toUser, byUser *entity.User) error {
	toUserId := toUser.UserID
	byUserId := byUser.UserID
	
	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		followDal := dal.Follow
		userDal := dal.User
		_, err := followDal.WithContext(ctx).Where(followDal.FollowbyID.Eq(toUserId), followDal.FollowerID.Eq(byUserId)).Delete()
		if err != nil {
			return err
		}
		_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(toUserId)).UpdateSimple(userDal.FollowerCount.Add(-1))
		if err != nil {
			return err
		}
		_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(byUserId)).UpdateSimple(userDal.FollowCount.Add(-1))
		if err != nil {
			return err
		}
		return nil
	})

	c.JSON(http.StatusOK, Response{StatusCode: 0})
	return err

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	
	userDal := dal.User
	viewUser, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(userId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	followDal := dal.Follow
	follows, err := followDal.WithContext(ctx).Where(followDal.FollowerID.Eq(userId)).Find()

	var usersController []User
	if len(follows) != 0 {
		for i := range follows {
			followbyUser, _ := userDal.WithContext(ctx).Where(userDal.UserID.Eq(follows[i].FollowbyID)).Take()
			usersController = append(
				usersController,
				*ConvertUserEntityToController(followbyUser, viewUser.UserID),
			)
		}
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: usersController,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	
	userDal := dal.User
	viewUser, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(userId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	followDal := dal.Follow
	follows, err := followDal.WithContext(ctx).Where(followDal.FollowbyID.Eq(userId)).Find()

	var usersController []User
	if len(follows) != 0 {
		for i := range follows {
			followerUser, _ := userDal.WithContext(ctx).Where(userDal.UserID.Eq(follows[i].FollowerID)).Take()
			usersController = append(
				usersController,
				*ConvertUserEntityToController(followerUser, viewUser.UserID),
			)
		}
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: usersController,
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
