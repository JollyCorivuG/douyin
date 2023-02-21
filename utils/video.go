package utils

import (
	"bytes"
	"douyin/config"
	"douyin/dao"
	"fmt"
	"log"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var (
	// 支持的视频格式
	VideoSupportFormat = map[string]struct{}{
		".mp4":  {},
		".mov":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".avi":  {},
	}
)

// 获取文件路径
func GetFileUrl(fileName string) string {
	return fmt.Sprintf("http://%s:%d/static/%s", config.Info.Ip, 8080, fileName)
}

// 根据用户id得到新的文件名
func NewFileName(userId int64) string {
	videoIndex, err := dao.DbMgr.QueryUserPublishVideoNum(userId)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%d-%d", userId, videoIndex)
}

// 截取视频的一帧为封面
func GetVideoCover(videoPath string, name string, frameNum int) string {
	// 新建一个缓冲区
	buf := bytes.NewBuffer(nil)

	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Println(err)
	}

	// 得到封面图
	coverImg, err := imaging.Decode(buf)
	if err != nil {
		log.Println(err)
	}

	coverName := name + ".jpeg"         // 图片名 (图片默认格式为jpg)
	savePath := "./static/" + coverName // 图片保存路径
	// 保存图片
	err = imaging.Save(coverImg, savePath)
	if err != nil {
		log.Println(err)

	}
	fmt.Println("debug dasdas", videoPath)
	return coverName
}
