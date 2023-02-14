package user_service

import (
	"douyin/model/example"
	"douyin/model/system"
)

const (
	like       = "1"
	cancelLike = "2"

	follow       = "1"
	cancelFollow = "2"
)

type UserServer interface {
	DoLogin(userName string, passWord string) (*example.Access, error)
	DoRegister(userName string, passWord string) (*example.Access, error)
	DoInfo(userId int64) (*system.UserInfo, error)
	DoLikeAction(userId int64, videoId int64, actionType string) error
	DoLikeList(userId int64) ([]*system.VideoInfo, error)
	DoFollowAction(followerId int64, followId int64, actionType string) error
	DoFollowList(userId int64) ([]*system.UserInfo, error)
	DoFollowerList(userId int64) ([]*system.UserInfo, error)
	DoFriendList(userId int64) ([]*system.UserInfo, error)
}

type server struct {
}

var Server UserServer

func init() {
	// 创建一个server对象
	Server = &server{}
}
