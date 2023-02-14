package user_handler

import (
	"douyin/model/common"
	"douyin/service/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LikeActionHandler(c *gin.Context) {
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

	// 得到videoId
	rawVideoId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(rawVideoId, 10, 64)

	// 得到actionType
	actionType := c.Query("action_type")

	// 调用服务
	if err := user_service.Server.DoLikeAction(userId, videoId, actionType); err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	})
}
