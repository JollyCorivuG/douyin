package user_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/system"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type friendListResponse struct {
	common.CommonResponse
	FriendList []*system.UserInfo `json:"user_list"`
}

func FriendListHandler(c *gin.Context) {
	rawUserId, ok1 := c.Get("user_id")
	userId, ok2 := rawUserId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, likeListResponse{
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
		c.JSON(http.StatusOK, likeListResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "user_id和token不一致",
			},
		})
		return
	}

	// 在数据库查询和userId互相关注的用户(即为好友)列表 (后期封装到service层)
	users, err := dao.DbMgr.QueryFriendUserByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, likeListResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 为users添加is_follow信息
	for index := range users {
		users[index].IsFollow = true
	}

	// 将数据返回到前端
	c.JSON(http.StatusOK, friendListResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "好友列表返回成功",
		},
		FriendList: users,
	})
}
