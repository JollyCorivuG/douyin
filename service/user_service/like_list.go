package user_service

import (
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type likeListFlow struct {
	userId int64
}

// 新建一个likeListFlow实例
func newLikeListFlow(userId int64) *likeListFlow {
	return &likeListFlow{userId: userId}
}

func (s *server) DoLikeList(userId int64) ([]*system.VideoInfo, error) {
	return newLikeListFlow(userId).do()
}

func (f *likeListFlow) do() ([]*system.VideoInfo, error) {
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
func (f *likeListFlow) checkPara() error {
	// 这里userId是上层解析来的, 一定合法
	return nil
}

func (f *likeListFlow) run(videos *[]*system.VideoInfo) error {
	videoList, err := dao.DbMgr.QueryLikeVideosByUserId(f.userId)
	if err != nil {
		return err
	}

	// 为videoList添加author和is_false信息
	for index := range videoList {
		videoList[index].VideoAuthor, err = dao.DbMgr.QueryUserByUserId(videoList[index].AuthorId)
		if err != nil {
			return err
		}
		videoList[index].IsFavorite = true
	}

	*videos = videoList

	return nil
}
