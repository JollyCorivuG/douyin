package video_service

import (
	"douyin/dao"
	"douyin/model/example"
	"log"
	"time"
)

// 包含handler层传来的参数
type feedFlow struct {
	latestTime time.Time
}

// 新建一个feedFlow实例
func newFeedFlow(latestTime time.Time) *feedFlow {
	return &feedFlow{latestTime: latestTime}
}

func (s *server) DoFeed(latestTime time.Time) (*example.VideoList, error) {
	return newFeedFlow(latestTime).do()
}

func (f *feedFlow) do() (*example.VideoList, error) {
	var videos *example.VideoList

	if err := f.checkPara(); err != nil {
		return nil, err
	}
	if err := f.run(&videos); err != nil {
		return nil, err
	}

	return videos, nil

}

// 检验参数
func (f *feedFlow) checkPara() error {
	// 传来的latestTime已经在上层经过解析, 一定合法
	return nil
}

func (f *feedFlow) run(videos **example.VideoList) error {
	videoList, err := dao.DbMgr.QueryVideosByLimit(maxVideoNum, f.latestTime)
	if err != nil {
		return err
	}

	// 将视频的作者附上
	for index := range videoList {
		author, err := dao.DbMgr.QueryUserByUserId(videoList[index].AuthorId)
		if err != nil {
			log.Println(err)
			continue
		}
		videoList[index].VideoAuthor = author
	}

	// 本次返回的视频中, 发布最早的时间, 作为下次请求时的latest_time
	nextTime := time.Now().Unix() / 1e6
	if videoSize := len(videoList); videoSize > 0 {
		nextTime = videoList[videoSize-1].CreateAt.Unix() / 1e6
	}

	*videos = &example.VideoList{
		Videos:   videoList,
		NextTime: nextTime,
	}

	return nil
}
