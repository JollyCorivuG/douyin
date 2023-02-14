package user_service

import (
	"douyin/cache"
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type followerListFlow struct {
	userId int64
}

// 新建一个followerListFlow实例
func newFollowerListFlow(userId int64) *followerListFlow {
	return &followerListFlow{userId: userId}
}

func (s *server) DoFollowerList(userId int64) ([]*system.UserInfo, error) {
	return newFollowerListFlow(userId).do()
}

func (f *followerListFlow) do() ([]*system.UserInfo, error) {
	var users []*system.UserInfo

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&users); err != nil {
		return nil, err
	}

	return users, nil
}

// 检验参数
func (f *followerListFlow) checkPara() error {
	// 这里userId是上层解析来的, 一定合法
	return nil
}

func (f *followerListFlow) run(users *[]*system.UserInfo) error {
	userList, err := dao.DbMgr.QueryFollowerUserByUserId(f.userId)
	if err != nil {
		return err
	}

	// 为userList添加is_follow信息
	for index := range userList {
		userList[index].IsFollow = cache.QueryUserIsFollowUser(f.userId, userList[index].UserId)
	}

	*users = userList

	return nil
}
