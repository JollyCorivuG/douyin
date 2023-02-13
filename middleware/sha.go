package middleware

import (
	"crypto/sha1"
	"douyin/model/common"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	minUserNameLen = 3
	maxUserNameLen = 32

	minPassWordLen = 5
	maxPassWordLen = 32
)

// 使用sha1对密码进行加密
func sha(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func ShaMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName := c.Query("username")
		passWord := c.Query("password")

		if len(userName) < minUserNameLen || len(userName) > maxUserNameLen || len(passWord) < minPassWordLen || len(passWord) > maxPassWordLen {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 403,
				StatusMsg:  "输入的用户名或密码不合法",
			})
			c.Abort()
			return
		}
		c.Set("password", sha(passWord))
		c.Next()
	}
}
