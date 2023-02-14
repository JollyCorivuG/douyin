package video_hander

import (
	"douyin/model/common"
	"douyin/model/system"
	"douyin/service/video_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type listResponse struct {
	common.CommonResponse
	VideoList []*system.VideoInfo `json:"video_list"`
}

func ListHandler(c *gin.Context) {
	rawId, ok1 := c.Get("user_id")
	userId, ok2 := rawId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, listResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "解析id出错",
			},
		})
		return
	}

	// user_id和token不一致
	userIdString := c.Query("user_id")
	userIdInt, _ := strconv.ParseInt(userIdString, 10, 64)
	if userId != userIdInt {
		c.JSON(http.StatusOK, listResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "user_id和token不一致",
			},
		})
		return
	}

	// 调用服务
	videos, err := video_service.Server.DoList(userId)
	if err != nil {
		c.JSON(http.StatusOK, listResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 数据返回给前端
	c.JSON(http.StatusOK, listResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "视频列表返回成功",
		},
		VideoList: videos,
	})
}
