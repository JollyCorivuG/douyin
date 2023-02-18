package message_service

import (
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type messageChatFlow struct {
	fromUserId int64
	toUserId   int64
	preMsgTime int64
}

// 新建一个messageChatFlow实例
func newMessageChatFlow(fromUserId int64, toUserId int64, preMsgTime int64) *messageChatFlow {
	return &messageChatFlow{fromUserId: fromUserId, toUserId: toUserId, preMsgTime: preMsgTime}
}

func (s *server) DoMessageChat(fromUserId int64, toUserId int64, preMsgTime int64) ([]*system.ChatMessage, error) {
	return newMessageChatFlow(fromUserId, toUserId, preMsgTime).do()
}

func (f *messageChatFlow) do() ([]*system.ChatMessage, error) {
	var messages []*system.ChatMessage

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&messages); err != nil {
		return nil, err
	}
	return messages, nil
}

// 检验参数
func (f *messageChatFlow) checkPara() error {
	// 参数一定合法
	return nil
}

func (f *messageChatFlow) run(messages *[]*system.ChatMessage) error {
	messageList, err := dao.DbMgr.QueryChatMessageByFromAndToUserId(f.fromUserId, f.toUserId, f.preMsgTime)
	if err != nil {
		return nil
	}

	*messages = messageList

	return nil
}
