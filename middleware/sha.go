package middleware

import (
	"crypto/sha1"
	"douyin/model/common"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用sha1对密码进行加密
func sha(s string) (string, error) {
	hash := sha1.New()
	if _, err := hash.Write([]byte(s)); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func ShaMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawPassWord := c.Query("password")

		passWord, err := sha(rawPassWord)
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 403,
				StatusMsg: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("password", passWord)
		c.Next()
	}
}
