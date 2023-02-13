package user_handler

import (
	"douyin/cache"
	"douyin/dao"
	"douyin/model/common"
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

	// 数据库更新操作 (后期封装到service层)
	if actionType == "1" {
		if cache.QueryUserIsLikeVideo(userId, videoId) {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "已经点赞,不能重复点赞",
			})
			return
		}
		if err := dao.DbMgr.UpdateVideoWhenLike(userId, videoId); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			})
			return
		}
	} else {
		if err := dao.DbMgr.UpdateVideoWhenCancelLike(userId, videoId); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			})
			return
		}
	}

	// redis缓存更新
	cache.UpdateVideoState(userId, videoId, actionType)

	c.JSON(http.StatusOK, common.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	})
}
