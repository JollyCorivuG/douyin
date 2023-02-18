package video_service

import (
	"douyin/dao"
	"douyin/model/system"
	"douyin/utils"
	"errors"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// 包含handler层传来的参数
type publishFlow struct {
	userId      int64
	title       string
	file        *multipart.FileHeader
	c           *gin.Context
	videoSuffix string
}

// 新建一个publishFlow实例
func newPublishFlow(userId int64, title string, file *multipart.FileHeader, c *gin.Context) *publishFlow {
	return &publishFlow{userId: userId, title: title, file: file, c: c}
}

func (s *server) DoPublish(userId int64, title string, file *multipart.FileHeader, c *gin.Context) error {
	return newPublishFlow(userId, title, file, c).do()
}

func (f *publishFlow) do() error {
	if err := f.checkPara(); err != nil {
		return err
	}
	if err := f.run(); err != nil {
		return err
	}

	return nil
}

// 检验参数
func (f *publishFlow) checkPara() error {
	// 检查视频格式
	f.videoSuffix = filepath.Ext(f.file.Filename)
	if _, ok := utils.VideoSupportFormat[f.videoSuffix]; !ok {
		return errors.New("不支持的视频格式")
	}
	return nil
}

func (f *publishFlow) run() error {
	// 根据用户id得到的文件名
	name := utils.NewFileName(f.userId)
	// 得到文视频名
	videoName := name + f.videoSuffix
	// 保存视频文件
	savePath := filepath.Join("./static", videoName)
	if err := f.c.SaveUploadedFile(f.file, savePath); err != nil {
		return err
	}
	// 截取视频的一帧作为封面
	coverName := utils.GetVideoCover("./static/"+videoName, name, 1)

	// 得到视频id和视频作者信息
	videoId, err := utils.GenerateId()
	if err != nil {
		return err
	}
	author, err := dao.DbMgr.QueryUserByUserId(f.userId)
	if err != nil {
		return err
	}

	videoInfo := &system.VideoInfo{
		VideoId:     videoId,
		AuthorId:    f.userId,
		VideoAuthor: author,
		PlayUrl:     utils.GetFileUrl(videoName),
		CoverUrl:    utils.GetFileUrl(coverName),
		Title:       f.title,
		CreateAt:    time.Now(),
	}
	if err := dao.DbMgr.AddVideoAndUpdateUser(videoInfo); err != nil {
		return err
	}

	return nil
}
