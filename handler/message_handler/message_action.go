package message_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/system"
	"douyin/utils"
	"net/http"
	"strconv"
	"time"

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

	// 得到参数后, 更新数据库 (后期封装到service层)
	if actionType == "1" {
		messageId, err := utils.GenerateId()
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "生成id出错",
			})
			return
		}
		chatMessage := &system.ChatMessage{
			MessageId:  messageId,
			ToUserId:   toUserId,
			FromUserId: userId,
			Content:    content,
			CreateTime: time.Now().Unix(),
		}
		if err := dao.DbMgr.AddChatMessage(chatMessage); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, common.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "发送信息成功",
	})
}
