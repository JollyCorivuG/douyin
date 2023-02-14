package video_service

import (
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type listFlow struct {
	userId int64
}

// 新建一个listFlow实例
func newListFlow(userId int64) *listFlow {
	return &listFlow{userId: userId}
}

func (s *server) DoList(userId int64) ([]*system.VideoInfo, error) {
	return newListFlow(userId).do()
}

func (f *listFlow) do() ([]*system.VideoInfo, error) {
	var videos []*system.VideoInfo

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&videos); err != nil {
		return nil, err
	}

	return videos, nil
}

// 检验参数
func (f *listFlow) checkPara() error {
	// 这里userId是上层解析来的, 一定合法
	return nil
}

func (f *listFlow) run(videos *[]*system.VideoInfo) error {
	videoList, err := dao.DbMgr.QueryVideosByUserId(f.userId)
	if err != nil {
		return err
	}

	// 作者信息
	author, err := dao.DbMgr.QueryUserByUserId(f.userId)
	if err != nil {
		return err
	}

	// 给返回的视频列表加上作者
	for index := range videoList {
		videoList[index].VideoAuthor = author
	}

	*videos = videoList

	return nil
}
