package user_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/example"
	"douyin/model/system"
	"douyin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册返回的响应
type registerResponse struct {
	common.CommonResponse
	*example.Access
}

// 处理注册
func RegisterHandler(c *gin.Context) {
	userName := c.Query("username")
	// passWord := c.Query("password")
	rawPassWord, ok1 := c.Get("password")
	passWord, ok2 := rawPassWord.(string)

	// 加密password时出错
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码加密时出错",
			},
		})
		return
	}

	// 用户名已经存在
	if dao.DbMgr.IsUserExistByUserName(userName) {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  "注册的用户名已经存在",
			},
		})
		return
	}

	// 创建一个用户
	userInfo := new(system.UserInfo)
	userInfo.UserName = userName
	userInfo.PassWord = passWord
	userId, err := utils.GenerateId()
	// 生成id时出错
	if err != nil {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	userInfo.UserId = userId

	tokenString, err := utils.ReleaseToken(userInfo.UserId)
	// 生成token时出错
	if err != nil {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 4,
				StatusMsg:  err.Error(),
			},
		})
	}

	// 注册成功
	c.JSON(http.StatusOK, registerResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		Access: &example.Access{
			UserId: userInfo.UserId,
			Token:  tokenString,
		},
	})

	// 将用户信息存入数据库
	dao.DbMgr.AddUser(userInfo)
}
