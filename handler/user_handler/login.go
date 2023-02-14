package user_handler

import (
	"douyin/model/common"
	"douyin/model/example"
	"douyin/service/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginResponse struct {
	common.CommonResponse
	*example.Access
}

func LoginHandler(c *gin.Context) {
	userName := c.Query("username")
	// passWord := c.Query("password")
	rawPassWord, ok1 := c.Get("password")
	passWord, ok2 := rawPassWord.(string)

	// 密码加密出错
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码加密时出错",
			},
		})
		return
	}

	// 调用服务
	access, err := user_service.Server.DoLogin(userName, passWord)
	if err != nil {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 登录成功
	c.JSON(http.StatusOK, loginResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		Access: access,
	})
}
