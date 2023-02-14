package video_service

import (
	"douyin/model/example"
	"douyin/model/system"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxVideoNum = 30
)

type VideoServer interface {
	DoFeed(latestTime time.Time) (*example.VideoList, error)
	DoList(userId int64) ([]*system.VideoInfo, error)
	DoPublish(userId int64, title string, file *multipart.FileHeader, c *gin.Context) error
}

type server struct {
}

var Server VideoServer

func init() {
	// 创建一个server对象
	Server = &server{}
}
