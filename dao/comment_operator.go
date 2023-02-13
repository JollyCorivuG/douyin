package dao

import (
	"douyin/model/system"
	"errors"
)

// 添加评论到数据库
func (dbMgr *manager) AddComment(commentInfo *system.CommentInfo) error {
	return dbMgr.DB.Create(commentInfo).Error
}

// 删除评论
func (dbMgr *manager) DeleteComment(commentId int64) error {
	return dbMgr.DB.Where("comment_id = ?", commentId).Delete(&system.CommentInfo{}).Error
}

// 根据id查找评论
func (dbMgr *manager) QueryCommentByCommentId(commentId int64) (*system.CommentInfo, error) {
	var commentInfo system.CommentInfo
	dbMgr.DB.Where("comment_id = ?", commentId).First(&commentInfo)
	if commentInfo.CommentId == 0 {
		return nil, errors.New("评论不存在")
	}
	return &commentInfo, nil
}

// 根据视频id查找该视频的评论
func (dbMgr *manager) QueryCommentByVideoId(videoId int64) ([]*system.CommentInfo, error) {
	var commentInfos []*system.CommentInfo
	if err := dbMgr.DB.Where("video_id = ?", videoId).Find(&commentInfos).Error; err != nil {
		return nil, err
	}
	return commentInfos, nil
}
