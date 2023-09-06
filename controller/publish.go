package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var static string = config.StaticUrl
var public string = config.PublicPath

// Publish videos
// Saved filename rule: username + create time + upload filename
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	userDal := dal.User
	videoDal := dal.Video
	
	user, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%s_%d_%s", user.Username, time.Now().Unix(), filename)
	saveFile := filepath.Join(public, finalName)

	video := entity.Video{
		// VideoID: NULL,
		UserID:     user.UserID,
		Token:      token,
		CreateTime: time.Now().Unix(),
		PlayURL:    static + finalName,
		// test cover
		CoverURL:      static + "bear.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         c.PostForm("title"),
	}

	err = dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		err := videoDal.WithContext(ctx).Create(&video)
		if err != nil {
			return err
		}
		_, err = userDal.WithContext(ctx).
			Where(userDal.UserID.Eq(user.UserID)).
			UpdateSimple(userDal.WorkCount.Add(1))
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	videoDal := dal.Video
	videos, err := videoDal.WithContext(ctx).Where(videoDal.UserID.Eq(userId)).Find()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User hasn't published videos"})
		return
	}

	var videosController []Video
	for i := range videos {
		videosController = append(
			videosController,
			*ConvertVideoEntityToController(videos[i]),
		)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videosController,
	})
}
