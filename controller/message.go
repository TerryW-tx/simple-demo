package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// "sync/atomic"
	"time"
	// "github.com/RaymondCode/simple-demo/model/dto"
	"github.com/RaymondCode/simple-demo/model/entity"
	"github.com/RaymondCode/simple-demo/dal"
)

var tempChat = map[string][]Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")

	userDal := dal.User
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

	messageDal := dal.Message
	message := entity.Message{
		UserID: fromUser.UserID,
		ToUserID: toUserId,
		MessageText: content,
		CreateTime: time.Now().Unix(),
		CreateDate: time.Now().Format("2006-01-02"),
	}

	err = messageDal.WithContext(ctx).Create(&message)
	if err == nil {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Message send failed"})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)

	userDal := dal.User
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
	
	messageDal := dal.Message
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

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
