package dao

import (
	"douyin/config"
	"douyin/model/system"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Manager interface {
	// 对用户的操作
	AddUser(userInfo *system.UserInfo) error
	IsUserExistByUserName(userName string) bool
	QueryUserByUserName(userName string) (*system.UserInfo, error)
	QueryUserByUserId(userId int64) (*system.UserInfo, error)
	QueryUserPublishVideoNum(userId int64) (int64, error)
	QueryUserIsLikeVideo(userId int64, videoId int64) bool
	UpdateUserWhenFollow(followerId int64, followId int64) error
	UpdateUserWhenCancelFollow(followerId int64, followId int64) error
	QueryFollowUserByUserId(userId int64) ([]*system.UserInfo, error)
	QueryFollowerUserByUserId(userId int64) ([]*system.UserInfo, error)
	QueryFriendUserByUserId(userId int64) ([]*system.UserInfo, error)

	// 对视频的操作
	AddVideoAndUpdateUser(videoInfo *system.VideoInfo) error
	QueryVideosByLimit(limitNum int, latestTime time.Time) ([]system.VideoInfo, error)
	QueryVideosByUserId(userId int64) ([]*system.VideoInfo, error)
	UpdateVideoWhenLike(userId int64, videoId int64) error
	UpdateVideoWhenCancelLike(userId int64, videoId int64) error
	QueryLikeVideosByUserId(userId int64) ([]*system.VideoInfo, error)

	// 对评论的操作
	AddComment(commentInfo *system.CommentInfo) error
	DeleteComment(commentId int64) error
	QueryCommentByCommentId(commentId int64) (*system.CommentInfo, error)
	QueryCommentByVideoId(videoId int64) ([]*system.CommentInfo, error)

	// 对信息的操作
	AddChatMessage(chatMessage *system.ChatMessage) error
	QueryChatMessageByFromAndToUserId(fromUserId int64, toUserId int64, preMsgTime int64) ([]*system.ChatMessage, error)
}

// 数据库管理者
var DbMgr Manager

type manager struct {
	DB *gorm.DB
}

func InitDataBase() {
	dsn := config.MysqlDbConnectString()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 取消默认事务
		PrepareStmt:            true, // 缓存预编译语句
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}, // 表名为单数
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建表
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&system.UserInfo{}, &system.VideoInfo{}, &system.UserLikeVideo{}, &system.UserRelation{}, &system.CommentInfo{}, &system.ChatMessage{})

	// 结构体实现接口
	DbMgr = &manager{db}
}
