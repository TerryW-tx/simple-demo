package controller

import (
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")

	userDal := dal.User
	videoDal := dal.Video

	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

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
		VideoID:     videoId,
		UserID:      userId,
		CommentText: c.Query("comment_text"),
		CreateTime:  time.Now().Unix(),
		CreateDate:  time.Now().Format("2006-01-02"),
	}

	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		commentDal := dal.Comment
		videoDal := dal.Video
		err := commentDal.WithContext(ctx).Create(&comment)
		if err != nil {
			return err
		}
		_, err = videoDal.WithContext(ctx).
			Where(videoDal.VideoID.Eq(videoId)).
			UpdateSimple(videoDal.CommentCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	})

	if err == nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 0},
			Comment:  *ConvertCommentEntityToController(&comment, user),
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
		_, err := commentDal.WithContext(ctx).
			Where(
				commentDal.VideoID.Eq(videoId),
				commentDal.UserID.Eq(userId),
			).Delete()
		if err != nil {
			return err
		}
		_, err = videoDal.WithContext(ctx).
			Where(videoDal.VideoID.Eq(videoId)).
			UpdateSimple(videoDal.CommentCount.Add(-1))
		if err != nil {
			return err
		}
		return nil
	})

	c.JSON(http.StatusOK, Response{StatusCode: 0})
	return err
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	userDal := dal.User
	videoDal := dal.Video

	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

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
