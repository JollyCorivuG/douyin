package comment_service

import (
	"douyin/dao"
	"douyin/model/system"
)

// 包含handler层传来的参数
type commentListFlow struct {
	userId  int64
	videoId int64
}

// 新建一个commentListFlow实例
func newCommentListFlow(userId int64, videoId int64) *commentListFlow {
	return &commentListFlow{userId: userId, videoId: videoId}
}

func (s *server) DoCommentList(userId int64, videoId int64) ([]*system.CommentInfo, error) {
	return newCommentListFlow(userId, videoId).do()
}

func (f *commentListFlow) do() ([]*system.CommentInfo, error) {
	var comments []*system.CommentInfo

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&comments); err != nil {
		return nil, err
	}

	return comments, nil
}

// 检验参数
func (f *commentListFlow) checkPara() error {
	// 参数一定合法
	return nil
}

func (f *commentListFlow) run(comments *[]*system.CommentInfo) error {
	commentList, err := dao.DbMgr.QueryCommentByVideoId(f.videoId)
	if err != nil {
		return err
	}

	// 为每条评论加上作者信息
	for index := range commentList {
		commentList[index].CommentPoster, err = dao.DbMgr.QueryUserByUserId(commentList[index].PosterId)
		if err != nil {
			return err
		}
	}

	*comments = commentList

	return nil
}
