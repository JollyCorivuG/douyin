package user_service

import (
	"douyin/dao"
	"douyin/model/example"
)

// 包含handler层传来的参数
type friendListFlow struct {
	userId int64
}

// 新建一个friendListFlow实例
func newFriendListFlow(userId int64) *friendListFlow {
	return &friendListFlow{userId: userId}
}

func (s *server) DoFriendList(userId int64) ([]*example.Friend, error) {
	return newFriendListFlow(userId).do()
}

func (f *friendListFlow) do() ([]*example.Friend, error) {
	var friends []*example.Friend

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&friends); err != nil {
		return nil, err
	}

	return friends, nil
}

// 检验参数
func (f *friendListFlow) checkPara() error {
	// 这里userId是上层解析来的, 一定合法
	return nil
}

func (f *friendListFlow) run(friends *[]*example.Friend) error {
	userList, err := dao.DbMgr.QueryFriendUserByUserId(f.userId)
	if err != nil {
		return err
	}

	// 为userList添加is_follow信息
	for index := range userList {
		userList[index].IsFollow = true
	}

	var friendList []*example.Friend

	// 加上最新信息
	for _, user := range userList {
		latestMsg, err := dao.DbMgr.QueryLatestMessageByFromAndToUserId(f.userId, user.UserId)
		if err != nil {
			return nil
		}
		friendList = append(friendList, &example.Friend{UserInfo: *user, LatestMessage: latestMsg})
	}

	*friends = friendList

	return nil
}
