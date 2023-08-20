package model

import (
  	"fmt"
  	// "time"
  	// "gorm.io/gorm"
)

type VideoInfo struct {
	VideoId		int64
	// VideoUrl	string
	UserId		int64
	Token		string
	CreateAt	int64
	PlayUrl		string
	CoverUrl	string
	FavoriteCount	int
	CommentCount	int
	Title		string
}

func CreateVideoInfo(v *VideoInfo) {
	fmt.Println("create video")
	db.Create(v)
}

func QueryLastNVideoInfo(time_stamp int64, video_num int) (vs []VideoInfo, e error) {
	var video_info_lists []VideoInfo
	err := db.Table("video_infos").Order("create_at DESC").Limit(video_num).Find(&video_info_lists).Error

	if err != nil {
		return nil, err
	}

	return video_info_lists, err
}
