package controller

import (
	"fmt"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
	"path/filepath"
	// "github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/RaymondCode/simple-demo/dal"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var static string = "http://120.55.103.230:8080/static/"
var videoIdSequence = int64(1)

func GenerateVideoId() int64 {
	atomic.AddInt64(&videoIdSequence, 1)
	return videoIdSequence
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
    	// for key, value := range c.Request.PostForm {
	// 	fmt.Printf("Key: %s, Value: %s\n", key, value)
    	// }

	userDal := dal.User
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
	saveFile := filepath.Join("./public/", finalName)
	
	fmt.Println("new video")
	video := entity.Video{
  		// VideoID: GenerateVideoId(),
		UserID: user.UserID,
        	Token: token,
        	CreateTime: time.Now().Unix(),
        	PlayURL: static + finalName,
        	CoverURL: static + "bear.jpg",
        	FavoriteCount: 0,
        	CommentCount: 0,
        	Title: c.PostForm("title"),
	}
	// model.CreateVideoInfo(&new_video)
	videoDal := dal.Video
	videoDal.WithContext(ctx).Create(&video)
	userDal.WithContext(ctx).Where(userDal.UserID.Eq(user.UserID)).UpdateSimple(userDal.WorkCount.Add(1))
	// saveFile := filepath.Join("./public/", 'test.mp4')
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

// PublishList all users have same publish video list
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
