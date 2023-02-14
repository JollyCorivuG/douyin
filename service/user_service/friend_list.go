package user_service

import (
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type friendListFlow struct {
	userId int64
}

// 新建一个friendListFlow实例
func newFriendListFlow(userId int64) *friendListFlow {
	return &friendListFlow{userId: userId}
}

func (s *server) DoFriendList(userId int64) ([]*system.UserInfo, error) {
	return newFriendListFlow(userId).do()
}

func (f *friendListFlow) do() ([]*system.UserInfo, error) {
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
func (f *friendListFlow) checkPara() error {
	// 这里userId是上层解析来的, 一定合法
	return nil
}

func (f *friendListFlow) run(users *[]*system.UserInfo) error {
	userList, err := dao.DbMgr.QueryFriendUserByUserId(f.userId)
	if err != nil {
		return err
	}

	// 为userList添加is_follow信息
	for index := range userList {
		userList[index].IsFollow = true
	}

	*users = userList

	return nil
}
