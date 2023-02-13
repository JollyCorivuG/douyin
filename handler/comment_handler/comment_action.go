package comment_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/system"
	"douyin/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type commenActionResponse struct {
	common.CommonResponse
	*system.CommentInfo
}

func CommentActionHandler(c *gin.Context) {
	rawUserId, ok1 := c.Get("user_id")
	userId, ok2 := rawUserId.(int64)
	if !ok1 || !ok2 {
		c.JSON(http.StatusOK, commenActionResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "解析id出错",
			},
		})
		return
	}

	videoIdString := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdString, 10, 64)
	actionType := c.Query("action_type")

	// 得到参数后, 更新数据库 (后期封装到service层)
	var commentInfo *system.CommentInfo
	if actionType == "1" {
		// 评论
		commentString := c.Query("comment_text")
		commentId, err := utils.GenerateId()
		if err != nil {
			c.JSON(http.StatusOK, commenActionResponse{
				CommonResponse: common.CommonResponse{
					StatusCode: 2,
					StatusMsg:  "生成id出错",
				},
			})
			return
		}
		commentInfo = &system.CommentInfo{
			CommentId:  commentId,
			PosterId:   userId,
			VideoId:    videoId,
			Content:    commentString,
			CreateDate: time.Now().Format("01-02"),
		}
		if err := dao.DbMgr.AddComment(commentInfo); err != nil {
			c.JSON(http.StatusOK, commenActionResponse{
				CommonResponse: common.CommonResponse{
					StatusCode: 3,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
	} else {
		// 删除评论
		commentIdString := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdString, 10, 64)
		commentInfo, _ = dao.DbMgr.QueryCommentByCommentId(commentId)
		if commentInfo.PosterId != userId {
			c.JSON(http.StatusOK, commenActionResponse{
				CommonResponse: common.CommonResponse{
					StatusCode: 2,
					StatusMsg:  "无法删除他人评论",
				},
			})
			return
		}
		
		if err := dao.DbMgr.DeleteComment(commentId); err != nil {
			c.JSON(http.StatusOK, commenActionResponse{
				CommonResponse: common.CommonResponse{
					StatusCode: 3,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
	}

	// 得到作者信息
	commentInfo.CommentPoster, _ = dao.DbMgr.QueryUserByUserId(commentInfo.PosterId)
	c.JSON(http.StatusOK, commenActionResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
		CommentInfo: commentInfo,
	})
}
