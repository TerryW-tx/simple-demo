package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/RaymondCode/simple-demo/model"
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
	query_time_stamp, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	fmt.Println(c.Query("latest_time"))
	
	video_infos, _ := model.QueryLastNVideoInfo(query_time_stamp, 3)
	if video_infos != nil {
		videos := VideoInfosToVideos(video_infos)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videos,
			NextTime:  time.Now().UnixNano(),
		})
	}
}

func VideoInfosToVideos(vis []model.VideoInfo) (vs []Video) {
	var video_lists []Video
	for i := range vis {
		video_lists = append(video_lists, Video{
			Id: vis[i].VideoId,
			Author: User{
				Id: 0,
				Name: "test_user",
				FollowCount: 0,
				FollowerCount: 0,
				IsFollow: false,
			},
			PlayUrl: vis[i].PlayUrl,
			CoverUrl: vis[i].CoverUrl,
			FavoriteCount: int64(vis[i].FavoriteCount),
			CommentCount: int64(vis[i].CommentCount),
			IsFavorite: false,
		})
	}
	return video_lists
	// fmt.Println(vis)
	// return vis
}
