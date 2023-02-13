package user_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/system"
	"net/http"

	"github.com/gin-gonic/gin"
)

type infoResponse struct {
	common.CommonResponse
	User *system.UserInfo `json:"user"`
}

func InfoHandler(c *gin.Context) {
	rawId, ok := c.Get("user_id")

	// id不存在
	if !ok {
		c.JSON(http.StatusOK, infoResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "id不存在",
			},
		})
		return
	}

	userId, ok := rawId.(int64)
	// 解析id出错
	if !ok {
		c.JSON(http.StatusOK, infoResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "解析id出错",
			},
		})
		return
	}

	userInfo, err := dao.DbMgr.QueryUserByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, infoResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 3,
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
		User: &system.UserInfo{
			UserId:   userInfo.UserId,
			UserName: userInfo.UserName,
		},
	})
}
