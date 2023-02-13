package user_handler

import (
	"douyin/cache"
	"douyin/dao"
	"douyin/model/common"
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

	if followId == followerId {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 2,
			StatusMsg:  "你不需要关注自己",
		})
		return
	}

	// 数据库更新操作 (后期封装到service层)
	if actionType == "1" {
		if cache.QueryUserIsFollowUser(followerId, followId) {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "已经关注,不能重复关注",
			})
			return
		}
		if err := dao.DbMgr.UpdateUserWhenFollow(followerId, followId); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			})
			return
		}
	} else {
		if err := dao.DbMgr.UpdateUserWhenCancelFollow(followerId, followId); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			})
			return
		}
	}

	// redis缓存更新
	cache.UpdateRelationState(followerId, followId, actionType)

	c.JSON(http.StatusOK, common.CommonResponse{
		StatusCode: 0,
		StatusMsg:  "操作成功",
	})
}
