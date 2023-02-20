package user_service

import (
	"douyin/dao"
	"douyin/model/example"
	"douyin/model/system"
	"douyin/utils"
	"errors"
)

// 包含handler层传来的参数
type registerFlow struct {
	userName string
	passWord string
}

// 新建一个registerFlow实例
func newRegisterFlow(userName string, passWord string) *registerFlow {
	return &registerFlow{userName: userName, passWord: passWord}
}

func (s *server) DoRegister(userName string, passWord string) (*example.Access, error) {
	return newRegisterFlow(userName, passWord).do()
}

func (f *registerFlow) do() (*example.Access, error) {
	var access *example.Access

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&access); err != nil {
		return nil, err
	}

	return access, nil
}

// 检验参数
func (f *registerFlow) checkPara() error {
	if dao.DbMgr.IsUserExistByUserName(f.userName) {
		return errors.New("注册的用户名已经存在")
	}
	return nil
}

func (f *registerFlow) run(access **example.Access) error {
	userId, err := utils.GenerateId()
	// 生成id时出错
	if err != nil {
		return err
	}

	tokenString, err := utils.ReleaseToken(userId)
	// 生成token出错
	if err != nil {
		return err
	}

	*access = &example.Access{
		UserId: userId,
		Token:  tokenString,
	}

	// 创建一个用户, 将其存入数据库
	userInfo := &system.UserInfo{
		UserName: f.userName,
		PassWord: f.passWord,
		UserId:   userId,
		Avatar:   utils.GetFileUrl("user_avatar.jpg"),
	}
	if err := dao.DbMgr.AddUser(userInfo); err != nil {
		return err
	}

	return nil
}
