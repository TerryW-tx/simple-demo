package controller

import (
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func FavoriteAction(c *gin.Context) {
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

	_, err = videoDal.WithContext(ctx).Where(videoDal.VideoID.Eq(videoId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User hasn't published videos"})
		return
	}

	if actionType == "1" {
		err = CreateFavorite(c, user.UserID, videoId)
	} else if actionType == "2" {
		err = CancelFavorite(c, user.UserID, videoId)
	}

	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

func CreateFavorite(c *gin.Context, userId, videoId int64) error {
	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		favorite := entity.Favorite{
			UserID:     userId,
			VideoID:    videoId,
			CreateTime: time.Now().Unix(),
		}
		favoriteDal := dal.Favorite
		videoDal := dal.Video
		err := favoriteDal.WithContext(ctx).Create(&favorite)
		if err != nil {
			return err
		}
		_, err = videoDal.WithContext(ctx).
			Where(videoDal.VideoID.Eq(videoId)).
			UpdateSimple(videoDal.FavoriteCount.Add(1))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func CancelFavorite(c *gin.Context, userId, videoId int64) error {
	err := dal.GetQueryByCtx(ctx).Transaction(func(tx *dal.Query) error {
		favoriteDal := dal.Favorite
		videoDal := dal.Video
		_, err := favoriteDal.WithContext(ctx).
			Where(
				favoriteDal.UserID.Eq(userId),
				favoriteDal.VideoID.Eq(videoId),
			).Delete()
		if err != nil {
			return err
		}
		_, err = videoDal.WithContext(ctx).
			Where(videoDal.VideoID.Eq(videoId)).
			UpdateSimple(videoDal.FavoriteCount.Add(-1))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func FavoriteList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	favoriteDal := dal.Favorite
	videoDal := dal.Video

	videos, err := videoDal.WithContext(ctx).Where(
		videoDal.Columns(videoDal.VideoID).In(
			favoriteDal.WithContext(ctx).
				Select(favoriteDal.VideoID).
				Where(favoriteDal.UserID.Eq(userId)),
		),
	).Find()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "No required videos in database"})
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
