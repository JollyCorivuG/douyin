package comment_handler

import (
	"douyin/dao"
	"douyin/model/common"
	"douyin/model/system"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentListResponse struct {
	common.CommonResponse
	CommentList []*system.CommentInfo `json:"comment_list"`
}

func CommentListHandler(c *gin.Context) {
	// rawUserId, ok1 := c.Get("user_id")
	// userId, ok2 := rawUserId.(int64)
	// if !ok1 || !ok2 {
	// 	c.JSON(http.StatusOK, commentListResponse{
	// 		CommonResponse: common.CommonResponse{
	// 			StatusCode: 1,
	// 			StatusMsg:  "解析id出错",
	// 		},
	// 	})
	// 	return
	// }

	videoIdString := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdString, 10, 64)
	// 在数据库根据videoId查询对应的评论 (后期封装到service层)
	comments, err := dao.DbMgr.QueryCommentByVideoId(videoId)
	if err != nil {
		c.JSON(http.StatusOK, commentListResponse{
			CommonResponse: common.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 为每条评论加上作者信息
	for index := range comments {
		comments[index].CommentPoster, err = dao.DbMgr.QueryUserByUserId(comments[index].PosterId)
		if err != nil {
			c.JSON(http.StatusOK, commentListResponse{
				CommonResponse: common.CommonResponse{
					StatusCode: 2,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, commentListResponse{
		CommonResponse: common.CommonResponse{
			StatusCode: 0,
			StatusMsg:  "评论列表返回成功",
		},
		CommentList: comments,
	})

}
