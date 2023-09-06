package controller

import (
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
// Feed video list
// Feed rule:
// 1. CreateTime < latest_time
// 2. if no more older videos exist, feed last several videos desc by CreateTime
func Feed(c *gin.Context) {
	var latestTime int64
	if c.Query("latest_time") == "" {
		latestTime = time.Now().Unix()
	} else {
		latestTime, _ = strconv.ParseInt(c.Query("latest_time"), 10, 64)
	}

	feedNum := 3
	videoDal := dal.Video
	videos, err := videoDal.WithContext(ctx).
		Limit(feedNum).
		Where(videoDal.CreateTime.Lte(latestTime)).
		Order(videoDal.CreateTime.Desc()).
		Find()
	if len(videos) == 0 {
		videos, err = videoDal.WithContext(ctx).
			Limit(feedNum).
			Order(videoDal.CreateTime.Desc()).
			Find()
	}
	if err != nil {
		return
	}

	var videosController []Video
	if len(videos) != 0 {
		for i := range videos {
			videosController = append(
				videosController,
				*ConvertVideoEntityToController(videos[i]),
			)
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videosController,
			NextTime:  videos[len(videos)-1].CreateTime,
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1},
			VideoList: videosController,
			NextTime:  time.Now().Unix(),
		})
	}
}
