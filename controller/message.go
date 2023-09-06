package controller

import (
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")

	userDal := dal.User
	messageDal := dal.Message

	fromUser, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(toUserId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	message := entity.Message{
		UserID:      fromUser.UserID,
		ToUserID:    toUserId,
		MessageText: content,
		CreateTime:  time.Now().Unix(),
		CreateDate:  time.Now().Format("2006-01-02"),
	}

	err = messageDal.WithContext(ctx).Create(&message)
	if err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Sending message failed"})
	}
}

func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	userDal := dal.User
	messageDal := dal.Message

	fromUser, err := userDal.WithContext(ctx).Where(userDal.Token.Eq(token)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	_, err = userDal.WithContext(ctx).Where(userDal.UserID.Eq(toUserId)).Take()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	messages, err := messageDal.WithContext(ctx).Where(
		messageDal.UserID.Eq(fromUser.UserID),
		messageDal.ToUserID.Eq(toUserId),
	).Or(
		messageDal.UserID.Eq(toUserId),
		messageDal.ToUserID.Eq(fromUser.UserID),
	).Find()

	var messagesController []Message
	if len(messages) != 0 {
		for i := range messages {
			messagesController = append(
				messagesController,
				*ConvertMessageEntityToController(messages[i]),
			)
		}
	}
	c.JSON(http.StatusOK, ChatResponse{
		Response:    Response{StatusCode: 0},
		MessageList: messagesController,
	})
}
