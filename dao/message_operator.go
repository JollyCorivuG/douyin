package dao

import "douyin/model/system"

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
