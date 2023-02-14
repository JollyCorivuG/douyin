package user_handler

import (
	"douyin/model/common"
	"douyin/service/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowActionHandler(c *gin.Context) {
	rawFollowerId, ok1 := c.Get("user_id")
	followerId, ok2 := rawFollowerId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "解析id出错",
		})
		return
	}

	followIdString := c.Query("to_user_id")
	followId, _ := strconv.ParseInt(followIdString, 10, 64)

	actionType := c.Query("action_type")

	// 调用服务
	if err := user_service.Server.DoFollowAction(followerId, followId, actionType); err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 3,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	})
}
