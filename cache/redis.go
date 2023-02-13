package cache

import (
	"context"
	"douyin/config"
	"fmt"

	"github.com/go-redis/redis/v8"
)

/*
	redis缓存 (存储映射关系)
	作用: 用于确认某个用户是否点赞某个视频或关注某个人
	对应关系: userId -> videoId, userId -> followId
*/

const (
	like     = "like"
	relation = "relation"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

func init() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Info.RedisDb.Host, config.Info.RedisDb.Port),
			Password: config.Info.RedisDb.PassWord,
			DB:       config.Info.RedisDb.Database,
		},
	)
}

// 当用户对视频点赞或取消点赞时更新状态 (actionType: 1-点赞，2-取消点赞)
func UpdateVideoState(userId int64, videoId int64, actionType string) {
	key := fmt.Sprintf("%s:%d", like, userId)

	// 点赞
	if actionType == "1" {
		rdb.SAdd(ctx, key, videoId)
		return
	}
	// 取消点赞
	rdb.SRem(ctx, key, videoId)
}

// 查询用户是否点赞该视频
func QueryUserIsLikeVideo(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", like, userId)
	return rdb.SIsMember(ctx, key, videoId).Val()
}

// 当用户关注某个用户时更新状态
func UpdateRelationState(followerId int64, followId int64, actionType string) {
	key := fmt.Sprintf("%s:%d", relation, followerId)
	// 关注
	if actionType == "1" {
		rdb.SAdd(ctx, key, followId)
		return
	}
	// 取消关注
	rdb.SRem(ctx, key, followId)
}

// 查询用户是否关注某个用户
func QueryUserIsFollowUser(followerId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", relation, followerId)
	return rdb.SIsMember(ctx, key, followId).Val()
}
