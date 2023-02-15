package message_handler

import (
	"douyin/model/common"
	"douyin/model/system"
	"douyin/service/message_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type messageChatResponse struct {
	common.CommonResponse
	messageList []*system.ChatMessage
}

func MessageChatHandler(c *gin.Context) {
	// 得到userId
	rawUserId, ok1 := c.Get("user_id")
	userId, ok2 := rawUserId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, messageChatResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "解析id出错",
			},
		})
		return
	}

	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)

	// 调用服务
	messages, err := message_service.Server.DoMessageChat(userId, toUserId)
	if err != nil {
		c.JSON(http.StatusOK, messageChatResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, messageChatResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "聊天记录显现成功",
		},
		messageList: messages,
	})

}
