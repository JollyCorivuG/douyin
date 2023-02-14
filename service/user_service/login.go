package user_service

import (
	"douyin/dao"
	"douyin/model/example"
	"douyin/utils"
	"errors"
)

// 包含handler层传来的参数
type loginFlow struct {
	userName string
	passWord string
	userId   int64
}

// 新建一个loginFlow实例
func newLoginFlow(userName string, passWord string) *loginFlow {
	return &loginFlow{userName: userName, passWord: passWord}
}

func (s *server) DoLogin(userName string, passWord string) (*example.Access, error) {
	return newLoginFlow(userName, passWord).do()
}

func (f *loginFlow) do() (*example.Access, error) {
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
func (f *loginFlow) checkPara() error {
	userInfo, err := dao.DbMgr.QueryUserByUserName(f.userName)
	// 用户名不存在
	if err != nil {
		return err
	}
	// 密码错误
	if f.passWord != userInfo.PassWord {
		return errors.New("密码错误")
	}

	f.userId = userInfo.UserId
	return nil
}

func (f *loginFlow) run(access **example.Access) error {
	tokenString, err := utils.ReleaseToken(f.userId)
	// 生成token出错
	if err != nil {
		return err
	}

	*access = &example.Access{
		UserId: f.userId,
		Token:  tokenString,
	}

	return nil
}
