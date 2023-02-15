package comment_handler

import (
	"douyin/model/common"
	"douyin/model/system"
	"douyin/service/comment_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commenActionResponse struct {
	common.CommonResponse
	*system.CommentInfo
}

func CommentActionHandler(c *gin.Context) {
	rawUserId, ok1 := c.Get("user_id")
	userId, ok2 := rawUserId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, commenActionResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "解析id出错",
			},
		})
		return
	}

	videoIdString := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdString, 10, 64)

	actionType := c.Query("action_type")

	// 调用服务
	comment, err := comment_service.Server.DoCommentAction(userId, videoId, actionType, c)
	if err != nil {
		c.JSON(http.StatusOK, commenActionResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, commenActionResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
		CommentInfo: comment,
	})
}
