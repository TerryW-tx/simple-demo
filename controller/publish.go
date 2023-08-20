package controller

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"github.com/RaymondCode/simple-demo/model"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

var static string = "http://120.55.103.230:8080/static/"

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if exist, _ := ValidateToken(token); !exist {
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
	// user := usersLoginInfo[token]
	user, _ := model.QueryUserInfoByToken(token)
	finalName := fmt.Sprintf("%s_%d_%s", user.UserName, time.Now().UnixNano(), filename)
	saveFile := filepath.Join("./public/", finalName)
	
	fmt.Println("new video")
	new_video := model.VideoInfo{
  		VideoId: 0,
		UserId: user.UserId,
        	Token: token,
        	CreateAt: time.Now().UnixNano(),
        	PlayUrl: static + finalName,
        	CoverUrl: static + "bear.jpg",
        	FavoriteCount: 0,
        	CommentCount: 0,
        	Title: filename,
	}
	model.CreateVideoInfo(&new_video)
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
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
