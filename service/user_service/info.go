package user_service

import (
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type infoFlow struct {
	userId int64
}

// 新建一个infoFlow实例
func newInfoFlow(userId int64) *infoFlow {
	return &infoFlow{userId: userId}
}

func (s *server) DoInfo(userId int64) (*system.UserInfo, error) {
	return newInfoFlow(userId).do()
}

func (f *infoFlow) do() (*system.UserInfo, error) {
	var user *system.UserInfo

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&user); err != nil {
		return nil, err
	}

	return user, nil
}

// 检验参数
func (f *infoFlow) checkPara() error {
	// 这里userId是上层解析来的, 一定合法
	return nil
}

func (f *infoFlow) run(user **system.UserInfo) error {
	userInfo, err := dao.DbMgr.QueryUserByUserId(f.userId)
	if err != nil {
		return err
	}

	*user = userInfo

	return nil
}
