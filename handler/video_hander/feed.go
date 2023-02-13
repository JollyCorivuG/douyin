package video_hander

import (
	"douyin/dao"
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

	if ok2 || rawTimeStamp != "" {
		// 如果得到了latest_time参数, 就将它解析为int64
		timeStamp, err := strconv.ParseInt(rawTimeStamp, 10, 64)
		if err != nil {
			log.Println(err)
		}
		latestTime = time.Unix(0, timeStamp*1e6) // 前端传来的时间戳是以ms为单位的
	}

	// 从数据库中得到视频, 并返回给前端 (后期封装到service层)
	videos, err := dao.DbMgr.QueryVideosByLimit(video_service.MaxVideoNum, latestTime)
	if err != nil {
		c.JSON(http.StatusOK, feedResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 将视频的作者附上
	for index := range videos {
		author, err := dao.DbMgr.QueryUserByUserId(videos[index].AuthorId)
		if err != nil {
			log.Println(err)
			continue
		}
		videos[index].VideoAuthor = author
	}

	// 本次返回的视频中, 发布最早的时间, 作为下次请求时的latest_time
	nextTime := time.Now().Unix() / 1e6
	if videoSize := len(videos); videoSize > 0 {
		nextTime = videos[videoSize-1].CreateAt.Unix() / 1e6
	}

	// 把响应返回给前端
	c.JSON(http.StatusOK, feedResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "推送视频成功",
		},
		VideoList: &example.VideoList{
			Videos:   videos,
			NextTime: nextTime,
		},
	})
}
