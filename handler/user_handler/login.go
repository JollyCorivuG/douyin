package user_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/example"
	"douyin/utils"
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

	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, registerResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码加密时出错",
			},
		})
		return
	}

	userInfo, err := dao.DbMgr.QueryUserByUserName(userName)
	// 用户不存在
	if err != nil {
		c.JSON(http.StatusOK, loginResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 密码错误
	if passWord != userInfo.PassWord {
		c.JSON(http.StatusOK, loginResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  "密码错误",
			},
		})
		return
	}

	tokenString, err := utils.ReleaseToken(userInfo.UserId)
	// 生成token时出错
	if err != nil {
		c.JSON(http.StatusOK, loginResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 4,
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
		Access: &example.Access{
			UserId: userInfo.UserId,
			Token:  tokenString,
		},
	})
}
