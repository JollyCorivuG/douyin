package example

import "douyin/model/system"

// 最新的一条消息
type LatestMessage struct {
	Content     string `json:"message"`
	MessageType int `json:"msg_type"`
}

// 返回的好友信息
type Friend struct {
	system.UserInfo 
	LatestMessage 
}
