package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/RaymondCode/simple-demo/dal"
	"net/http"
	"time"
	"fmt"
	"strconv"
)

// test

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	fmt.Println(c.Query("latest_time"))

	videoDal := dal.Video
	videos, _ := videoDal.WithContext(ctx).Limit(3).Where(videoDal.CreateTime.Lte(latestTime)).Order(videoDal.CreateTime.Desc()).Find()
	// video_infos, _ := model.QueryLastNVideoInfo(query_time_stamp, 3)
	if videos != nil {
		var videosController []Video
		for i := range videos {
			videosController = append(
				videosController, 
				ConvertVideoEntityToController(videos[i]),
			)
		}
		// videos := VideoInfosToVideos(video_infos)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videosController,
			NextTime:  videos[len(videos)-1].CreateTime,
		})
	}
}

func ConvertVideoEntityToController(video *entity.Video) *Video {
	userDal := dal.User
	user, _ := userDal.WithContext(ctx).Where(userDal.UserID.Eq(video.UserID)).Take()
	videoController := controller.video{
		Id: video.VideoId,
		Author: user,
		PlayUrl: video.PlayURL,
		CoverUrl: video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount: video.CommentCount,
		IsFavorite: false,
	}
}
