package message_service

import "douyin/model/system"

const (
	postMessage = "1"
)

type MessageServer interface {
	DoMessageAction(fromUserId int64, toUserId int64, content string, actionType string) error
	DoMessageChat(fromUserId int64, toUserId int64) ([]*system.ChatMessage, error)
}

type server struct {
}

var Server MessageServer

func init() {
	// 创建一个server对象
	Server = &server{}
}
