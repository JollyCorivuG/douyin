package video_hander

import (
	"douyin/model/common"
	"douyin/service/video_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishHandler(c *gin.Context) {
	rawId, ok1 := c.Get("user_id")
	userId, ok2 := rawId.(int64)

	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "解析id出错",
		})
		return
	}

	title := c.PostForm("title")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 多文件上传
	files := form.File["data"]
	for _, file := range files {
		if err := video_service.Server.DoPublish(userId, title, file, c); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			})
			return
		}

		// 上传成功
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "上传成功",
		})
	}
}
