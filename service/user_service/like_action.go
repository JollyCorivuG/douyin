package user_service

import (
	"douyin/cache"
	"douyin/dao"
	"errors"
)

// 包含handler层传来的参数
type likeActionFlow struct {
	userId     int64
	videoId    int64
	actionType string
}

// 新建一个likeActionFlow实例
func newLikeActionFlow(userId int64, videoId int64, actionType string) *likeActionFlow {
	return &likeActionFlow{userId: userId, videoId: videoId, actionType: actionType}
}

func (s *server) DoLikeAction(userId int64, videoId int64, actionType string) error {
	return newLikeActionFlow(userId, videoId, actionType).do()
}

func (f *likeActionFlow) do() error {
	if err := f.checkPara(); err != nil {
		return err
	}
	if err := f.run(); err != nil {
		return err
	}

	return nil
}

// 检验参数
func (f *likeActionFlow) checkPara() error {
	if f.actionType == like && cache.QueryUserIsLikeVideo(f.userId, f.videoId) {
		return errors.New("已经点赞, 不能重复点赞")
	}
	return nil
}

func (f *likeActionFlow) run() error {
	if f.actionType == like {
		if err := dao.DbMgr.UpdateVideoWhenLike(f.userId, f.videoId); err != nil {
			return err
		}
	} else {
		if err := dao.DbMgr.UpdateVideoWhenCancelLike(f.userId, f.videoId); err != nil {
			return err
		}
	}

	// redis缓存更新
	cache.UpdateVideoState(f.userId, f.videoId, f.actionType)

	return nil
}
