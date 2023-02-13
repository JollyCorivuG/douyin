package video_hander

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/system"
	"douyin/utils"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func PublishHandler(c *gin.Context) {
	rawId, ok1 := c.Get("user_id")
	userId, ok2 := rawId.(int64)

	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 1,
			StatusMsg:  "解析id出错",
		})
		return
	}

	title := c.PostForm("title")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
	}

	// 多文件上传
	files := form.File["data"]
	for _, file := range files {
		videoSuffix := filepath.Ext(file.Filename)
		if _, ok := utils.VideoSupportFormat[videoSuffix]; !ok {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 3,
				StatusMsg:  "不支持的视频格式",
			})
			continue
		}

		// 根据用户id得到的文件名
		name := utils.NewFileName(userId)
		// 得到文视频名
		videoName := name + videoSuffix
		// 保存视频文件
		savePath := filepath.Join("./static", videoName)
		err := c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 4,
				StatusMsg:  err.Error(),
			})
			continue
		}
		// 截取视频的一帧作为封面
		coverName := utils.GetVideoCover("./static/"+videoName, name, 1)

		// 将视频保存到数据库 (后期封装到service层)
		videoInfo := new(system.VideoInfo)
		videoId, err := utils.GenerateId()
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 5,
				StatusMsg:  err.Error(),
			})
			continue
		}
		videoInfo.VideoId = videoId
		rawAuthor, err := dao.DbMgr.QueryUserByUserId(userId)
		if err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 6,
				StatusMsg:  err.Error(),
			})
			continue
		}
		videoInfo.AuthorId = userId
		videoInfo.VideoAuthor = rawAuthor
		videoInfo.PlayUrl = utils.GetFileUrl(videoName)
		videoInfo.CoverUrl = utils.GetFileUrl(coverName)
		videoInfo.Title = title
		videoInfo.CreateAt = time.Now()
		if err := dao.DbMgr.AddVideo(videoInfo); err != nil {
			c.JSON(http.StatusOK, common.CommonResponse{
				StatusCode: 7,
				StatusMsg:  err.Error(),
			})
		}

		// 上传成功
		c.JSON(http.StatusOK, common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  videoName + "上传成功",
		})
	}
}
