package comment_service

import (
	"douyin/model/system"

	"github.com/gin-gonic/gin"
)

const (
	doComment       = "1"
	cancelComment = "2"
)

type CommentServer interface {
	DoCommentAction(userId int64, videoId int64, actionType string, c *gin.Context) (*system.CommentInfo, error)
	DoCommentList(userId int64, videoId int64) ([]*system.CommentInfo, error)
}

type server struct {
}

var Server CommentServer

func init() {
	// 创建一个server对象
	Server = &server{}
}
