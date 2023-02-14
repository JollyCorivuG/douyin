package user_handler

import (
	"douyin/model/common"
	"douyin/model/system"
	"douyin/service/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type infoResponse struct {
	common.CommonResponse
	User *system.UserInfo `json:"user"`
}

func InfoHandler(c *gin.Context) {
	rawUserId, ok1 := c.Get("user_id")
	userId, ok2 := rawUserId.(int64)

	// 解析id出错
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, infoResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "解析id出错",
			},
		})
		return
	}

	// 调用服务
	user, err := user_service.Server.DoInfo(userId)
	if err != nil {
		c.JSON(http.StatusOK, infoResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, infoResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "获取用户信息成功",
		},
		User: user,
	})
}
