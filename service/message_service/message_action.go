package message_service

import (
	"douyin/dao"
	"douyin/model/system"
	"douyin/utils"
	"time"
)

// 包含handler层传来的参数
type messageActionFlow struct {
	fromUserId int64
	toUserId   int64
	content    string
	actionType string
}

// 新建一个messageActionFlow实例
func newMessageActionFlow(fromUserId int64, toUserId int64, content string, actionType string) *messageActionFlow {
	return &messageActionFlow{fromUserId: fromUserId, toUserId: toUserId, content: content, actionType: actionType}
}

func (s *server) DoMessageAction(fromUserId int64, toUserId int64, content string, actionType string) error {
	return newMessageActionFlow(fromUserId, toUserId, content, actionType).do()
}

func (f *messageActionFlow) do() error {
	if err := f.checkPara(); err != nil {
		return err
	}
	if err := f.run(); err != nil {
		return err
	}

	return nil
}

// 检验参数
func (f *messageActionFlow) checkPara() error {
	// 参数一定合法
	return nil
}

func (f *messageActionFlow) run() error {
	if f.actionType == postMessage {
		messageId, err := utils.GenerateId()
		if err != nil {
			return err
		}

		chatMessage := &system.ChatMessage{
			MessageId:  messageId,
			ToUserId:   f.toUserId,
			FromUserId: f.fromUserId,
			Content:    f.content,
			CreateTime: time.Now().Unix(),
		}
		if err := dao.DbMgr.AddChatMessage(chatMessage); err != nil {
			return err
		}
	}
	return nil
}
