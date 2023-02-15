package comment_service

import (
	"douyin/dao"
	"douyin/model/system"
	"douyin/utils"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 包含handler层传来的参数
type commentActionFlow struct {
	userId     int64
	videoId    int64
	actionType string
	c          *gin.Context
}

// 新建一个commentActionFlow实例
func newCommentActionFlow(userId int64, videoId int64, actionType string, c *gin.Context) *commentActionFlow {
	return &commentActionFlow{userId: userId, videoId: videoId, actionType: actionType, c: c}
}

func (s *server) DoCommentAction(userId int64, videoId int64, actionType string, c *gin.Context) (*system.CommentInfo, error) {
	return newCommentActionFlow(userId, videoId, actionType, c).do()
}

func (f *commentActionFlow) do() (*system.CommentInfo, error) {
	var comment *system.CommentInfo

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&comment); err != nil {
		return nil, err
	}

	return comment, nil
}

// 检验参数
func (f *commentActionFlow) checkPara() error {
	// 参数一定合法
	return nil
}

func (f *commentActionFlow) run(comment **system.CommentInfo) error {
	if f.actionType == doComment {
		commentString := f.c.Query("comment_text")
		commentId, err := utils.GenerateId()
		if err != nil {
			return err
		}

		commentInfo := &system.CommentInfo{
			CommentId:  commentId,
			PosterId:   f.userId,
			VideoId:    f.videoId,
			Content:    commentString,
			CreateDate: time.Now().Format("01-02"),
		}
		if err := dao.DbMgr.AddComment(commentInfo); err != nil {
			return err
		}

		*comment = commentInfo

	} else {
		commentIdString := f.c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdString, 10, 64)
		commentInfo, _ := dao.DbMgr.QueryCommentByCommentId(commentId)
		if commentInfo.PosterId != f.userId {
			return errors.New("无法删除他人评论")
		}

		if err := dao.DbMgr.DeleteComment(commentId); err != nil {
			return err
		}

		*comment = commentInfo
	}

	(*comment).CommentPoster, _ = dao.DbMgr.QueryUserByUserId((*comment).PosterId)

	return nil
}
