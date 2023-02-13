package middleware

import (
	"douyin/model/common"
	"douyin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")
		if tokenString == "" {
			tokenString = c.PostForm("token")
		}

		// 用户不存在
		if tokenString == "" {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 401,
				StatusMsg:  "用户不存在",
			})
			c.Abort()
			return
		}

		cliams, err := utils.ParseToken(tokenString)
		// 解析token出错
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 403,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}

		// token过期
		if time.Now().Unix() > cliams.ExpiresAt {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort()
			return
		}

		c.Set("user_id", cliams.UserId)
		c.Next()
	}
}
