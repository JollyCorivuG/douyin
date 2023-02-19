package dao

import (
	"douyin/model/example"
	"douyin/model/system"
)

// 添加聊天信息到数据库
func (dbMgr *manager) AddChatMessage(chatMessage *system.ChatMessage) error {
	return dbMgr.DB.Create(chatMessage).Error
}

// 根据发送者和接收者id查询消息
func (dbMgr *manager) QueryChatMessageByFromAndToUserId(fromUserId int64, toUserId int64, preMsgTime int64) ([]*system.ChatMessage, error) {
	var messageList []*system.ChatMessage
	if err := dbMgr.DB.Where("from_user_id = ? AND to_user_id = ? OR from_user_id = ? AND to_user_id = ?", fromUserId, toUserId, toUserId, fromUserId).Where("create_time > ?", preMsgTime).Find(&messageList).Error; err != nil {
		return nil, err
	}
	return messageList, nil
}

// 根据发送者和接收者id查询最新的一条消息
func (dbMgr *manager) QueryLatestMessageByFromAndToUserId(fromUserId int64, toUserId int64) (example.LatestMessage, error) {
	var latestMsg system.ChatMessage
	if err := dbMgr.DB.Where("from_user_id = ? AND to_user_id = ? OR from_user_id = ? AND to_user_id = ?", fromUserId, toUserId, toUserId, fromUserId).Last(&latestMsg).Error; err != nil {
		return example.LatestMessage{}, err
	}

	if latestMsg.FromUserId == fromUserId {
		return example.LatestMessage{Content: latestMsg.Content, MessageType: 0}, nil
	}
	return example.LatestMessage{Content: latestMsg.Content, MessageType: 1}, nil
}
