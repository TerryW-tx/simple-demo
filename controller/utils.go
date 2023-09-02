
package controller

import (
	// "fmt"
	// "time"
	// "context"
	"github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/RaymondCode/simple-demo/dal"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "sync/atomic"
	// "strconv"
)

func ConvertUserEntityToDto(user *entity.User) *dto.User {
	userDto := dto.User{
		UserID: user.UserID,
		Username: user.Username,
		Password: user.Password,
		CreateTime: user.CreateTime,
		Token: user.Token,
		TokenUpdateTime: user.TokenUpdateTime,
		Avatar: user.Avatar,
		BackgroundImage: user.BackgroundImage,
		FollowCount: int(user.FollowCount),
		FollowerCount: int(user.FollowerCount),
		FavoriteCount: int(user.FavoriteCount),
		WorkCount: int(user.WorkCount),
	}
	return &userDto
}

func ConvertUserEntityToController(user *entity.User, followerId int64) *User {	
	followDal := dal.Follow
	_, followErr := followDal.WithContext(ctx).Where(
		followDal.FollowbyID.Eq(user.UserID),
		followDal.FollowerID.Eq(followerId),
	).Take()
	userController := User{
		Id: user.UserID,
		Name: user.Username,
		FollowCount: user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow: followErr == nil,
	}
	return &userController
}

func ConvertVideoEntityToController(video *entity.Video) *Video {
	userDal := dal.User
	user, _ := userDal.WithContext(ctx).Where(userDal.UserID.Eq(video.UserID)).Take()
	// followerId := strconv.ParseInt(video.UserID, 10, 64)
	videoController := Video{
		Id: video.VideoID,
		Author: *ConvertUserEntityToController(user, video.UserID),
		PlayUrl: video.PlayURL,
		CoverUrl: video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount: video.CommentCount,
		IsFavorite: false,
	}
	return &videoController
}


func ConvertCommentEntityToController(comment *entity.Comment, user *entity.User) *Comment {
	commentController := Comment{
		Id: comment.CommentID,
		User: *ConvertUserEntityToController(user, comment.UserID),
		Content: comment.CommentText,
		CreateDate: comment.CreateDate,
	}
	return &commentController
}
