package user_service

import (
	"douyin/cache"
	"douyin/dao"
	"errors"
)

// 包含handler层传来的参数
type followActionFlow struct {
	followerId int64
	followId   int64
	actionType string
}

// 新建一个followActionFlow实例
func newFollowActionFlow(followerId int64, followId int64, actionType string) *followActionFlow {
	return &followActionFlow{followerId: followerId, followId: followId, actionType: actionType}
}

func (s *server) DoFollowAction(followerId int64, followId int64, actionType string) error {
	return newFollowActionFlow(followerId, followId, actionType).do()
}

func (f *followActionFlow) do() error {
	if err := f.checkPara(); err != nil {
		return err
	}
	if err := f.run(); err != nil {
		return err
	}

	return nil
}

// 检验参数
func (f *followActionFlow) checkPara() error {
	if f.followId == f.followerId {
		return errors.New("你不需要关注自己")
	}
	if f.actionType == follow && cache.QueryUserIsFollowUser(f.followerId, f.followId) {
		return errors.New("已经关注, 不能重复关注")
	}
	return nil
}

func (f *followActionFlow) run() error {
	if f.actionType == follow {
		if err := dao.DbMgr.UpdateUserWhenFollow(f.followerId, f.followId); err != nil {
			return err
		}
	} else {
		if err := dao.DbMgr.UpdateUserWhenFollow(f.followerId, f.followId); err != nil {
			return err
		}
	}

	// redis缓存更新
	cache.UpdateRelationState(f.followerId, f.followId, f.actionType)

	return nil
}
