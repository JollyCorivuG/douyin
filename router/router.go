package router

import (
	"douyin/handler/comment_handler"
	"douyin/handler/message_handler"
	"douyin/handler/user_handler"
	"douyin/handler/video_hander"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")

	rootGroup := router.Group("/douyin")

	// 用户路由
	userGroup := rootGroup.Group("/user")
	{
		userGroup.GET("/", middleware.JwtMiddleWare(), user_handler.InfoHandler)
		userGroup.POST("/register/", middleware.ShaMiddleWare(), user_handler.RegisterHandler)
		userGroup.POST("/login/", middleware.ShaMiddleWare(), user_handler.LoginHandler)
	}

	// feed路由
	feedGroup := rootGroup.Group("/feed")
	{
		feedGroup.GET("/", middleware.VideoFeedMiddleWare(), video_hander.FeedHandler)
	}

	// publish路由
	publishGroup := rootGroup.Group("/publish")
	{
		publishGroup.POST("/action/", middleware.JwtMiddleWare(), video_hander.PublishHandler)
		publishGroup.GET("/list/", middleware.JwtMiddleWare(), video_hander.ListHandler)
	}

	// favorite路由
	favoriteGroup := rootGroup.Group("/favorite")
	{
		favoriteGroup.POST("/action/", middleware.JwtMiddleWare(), user_handler.LikeActionHandler)
		favoriteGroup.GET("/list/", middleware.JwtMiddleWare(), user_handler.LikeListHandler)
	}

	// comment路由
	commentGroup := rootGroup.Group("/comment")
	{
		commentGroup.POST("/action/", middleware.JwtMiddleWare(), comment_handler.CommentActionHandler)
		commentGroup.GET("/list/", middleware.JwtMiddleWare(), comment_handler.CommentListHandler)
	}

	// relation路由
	relationGroup := rootGroup.Group("/relation")
	{
		relationGroup.POST("/action/", middleware.JwtMiddleWare(), user_handler.FollowActionHandler)
		relationGroup.GET("/follow/list/", middleware.JwtMiddleWare(), user_handler.FollowListHandler)
		relationGroup.GET("/follower/list/", middleware.JwtMiddleWare(), user_handler.FollowerListHandler)
		relationGroup.GET("/friend/list/", middleware.JwtMiddleWare(), user_handler.FriendListHandler)
	}

	// message路由
	messageGroup := rootGroup.Group("/message")
	{
		messageGroup.POST("/action/", middleware.JwtMiddleWare(), message_handler.MessageActionHandler)
		messageGroup.GET("/chat/", middleware.JwtMiddleWare(), message_handler.MessageChatHandler)
	}

	return router
}
