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

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	
	userDal := dal.User
	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videoDal := dal.Video
	video, err := videoDal.WithContext(ctx).Where(videoDal.VideoID.Eq(videoId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User hasn't published videos"})
		return
	}
	
	if actionType == "1" {
		err = CreateComment(c, user, video)
	} else if actionType == "2" {
		err = CancelComment(c, user, video)
	}
}

func CreateComment(c *gin.Context, user *entity.User, video *entity.Video) error {
	userId := user.UserID
	videoId := video.VideoID
	
	comment := entity.Comment{
		VideoID: videoId,
		UserID: userId,
		CommentText: c.Query("comment_text"),
		CreateTime: time.Now().Unix(),
		CreateDate: time.Now().Format("2006-01-02"),
	}

	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		// comment := entity.Comment{
		// 	VideoID: videoId,
		// 	UserID: userId,
		// 	CommentText: c.Query("comment_text"),
		// 	CreateTime: time.Now().Unix(),
		// 	CreateDate: time.Now().Format("2006-01-02"),
		// }
		commentDal := dal.Comment
		videoDal := dal.Video
		err := commentDal.WithContext(ctx).Create(&comment)
		if err != nil {
			return err
		}
		_, err = videoDal.WithContext(ctx).Where(videoDal.VideoID.Eq(videoId)).UpdateSimple(videoDal.CommentCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	})

	if err == nil {
		// commentDal := dal.Comment
		// comment, err := commentDal.WithContext(ctx).Where(
		// 	commentDal.UserID.Eq(userId),
		// 	commentDal.VideoID.Eq(videoId),
		// ).Take()

		// if err != nil {
		// 	return err
		// }

		// followDal := dal.Follow
		// _, followErr := followDal.WithContext(ctx).Where(
		// 	followDal.FollowbyID.Eq(video.UserID),
		// 	followDal.FollowerID.Eq(user.UserID),
		// ).Take()
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 0},
			Comment: *ConvertCommentEntityToController(&comment, user),
		})
	}

	return err
}


func CancelComment(c *gin.Context, user *entity.User, video *entity.Video) error {
	userId := user.UserID
	videoId := video.VideoID
	
	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		commentDal := dal.Comment
		videoDal := dal.Video
		_, err := commentDal.WithContext(ctx).Where(commentDal.VideoID.Eq(videoId), commentDal.UserID.Eq(userId)).Delete()
		if err != nil {
			return err
		}
		_, err = videoDal.WithContext(ctx).Where(videoDal.VideoID.Eq(videoId)).UpdateSimple(videoDal.CommentCount.Add(-1))
		if err != nil {
			return err
		}
		return nil
	})

	c.JSON(http.StatusOK, Response{StatusCode: 0})
	return err
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	userDal := dal.User
	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videoDal := dal.Video
	_, err = videoDal.WithContext(ctx).Where(videoDal.VideoID.Eq(videoId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User hasn't published videos"})
		return
	}

	commentDal := dal.Comment
	comments, err := commentDal.WithContext(ctx).Where(commentDal.VideoID.Eq(videoId)).Find()

	var commentsController []Comment
	if len(comments) != 0 {
		for i := range comments {
			commentsController = append(
				commentsController,
				*ConvertCommentEntityToController(comments[i], user),
			)
		}
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: commentsController,
	})
}
