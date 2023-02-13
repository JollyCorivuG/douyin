package middleware

import (
	"douyin/model/common"
	"douyin/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 视频推送中间件
func VideoFeedMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, ok := c.GetQuery("token")

		// 未登录状态
		if !ok {
			c.Next()
			return
		}

		// 登录状态
		cliams, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 403,
				StatusMsg:  err.Error(),
			})
		}

		// token过期
		if time.Now().Unix() > cliams.ExpiresAt {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
		}

		c.Set("user_id", cliams.UserId)
		c.Next()
	}
}
