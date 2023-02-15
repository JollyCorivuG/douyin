package message_handler

import (
	"douyin/model/common"
	"douyin/service/message_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MessageActionHandler(c *gin.Context) {
	// 得到userId
	rawUserId, ok1 := c.Get("user_id")
	userId, ok2 := rawUserId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "解析id出错",
		})
		return
	}

	toUserIdString := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64)

	actionType := c.Query("action_type")

	content := c.Query("content")

	// 调用服务
	if err := message_service.Server.DoMessageAction(userId, toUserId, content, actionType); err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "发送信息成功",
	})
}
