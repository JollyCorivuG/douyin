package video_hander

import (
	"douyin/model/common"
	"douyin/model/example"
	"douyin/service/video_service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type feedResponse struct {
	common.CommonResponse
	*example.VideoList
}

func FeedHandler(c *gin.Context) {
	// userId, ok1 := c.Get("user_id") // 得到登录用户的id
	rawTimeStamp, ok2 := c.GetQuery("latest_time")
	latestTime := time.Now()

	if ok2 || (rawTimeStamp != "" && rawTimeStamp != "0") {
		// 如果得到了latest_time参数, 就将它解析为int64
		timeStamp, err := strconv.ParseInt(rawTimeStamp, 10, 64)
		if err != nil {
			log.Println(err)
		}
		latestTime = time.Unix(0, timeStamp*1e6) // 前端传来的时间戳是以ms为单位的
	}

	// 调用服务
	videos, err := video_service.Server.DoFeed(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, feedResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	// 把响应返回给前端
	c.JSON(http.StatusOK, feedResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "推送视频成功",
		},
		VideoList: videos,
	})
}
