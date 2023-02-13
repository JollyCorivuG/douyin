package dao

import (
	"douyin/cache"
	"douyin/model/system"
	"errors"
	"log"

	"gorm.io/gorm"
)

// 添加用户到数据库
func (dbMgr *manager) AddUser(userInfo *system.UserInfo) error {
	return dbMgr.DB.Create(userInfo).Error
}

// 根据用户名查询用户是否存在
func (dbMgr *manager) IsUserExistByUserName(userName string) bool {
	var userInfos []system.UserInfo
	if err := dbMgr.DB.Where("user_name = ?", userName).Find(&userInfos).Error; err != nil {
		log.Println(err)
	}

	if len(userInfos) > 0 {
		return true
	}
	return false
}

// 根据用户名查询用户
func (dbMgr *manager) QueryUserByUserName(userName string) (*system.UserInfo, error) {
	var userInfo system.UserInfo
	dbMgr.DB.Where("user_name = ?", userName).First(&userInfo)
	if userInfo.UserId == 0 {
		return nil, errors.New("用户不存在")
	}
	return &userInfo, nil
}

// 根据用户id查询用户
func (dbMgr *manager) QueryUserByUserId(userId int64) (*system.UserInfo, error) {
	var userInfo system.UserInfo
	dbMgr.DB.Where("user_id = ?", userId).First(&userInfo)
	if userInfo.UserId == 0 {
		return nil, errors.New("用户不存在")
	}
	return &userInfo, nil
}

// 查询用户发布视频的数量
func (dbMgr *manager) QueryUserPublishVideoNum(userId int64) (int64, error) {
	var videoCount int64
	if err := dbMgr.DB.Model(&system.VideoInfo{}).Where("author_id = ?", userId).Count(&videoCount).Error; err != nil {
		return 0, err
	}
	return videoCount, nil
}

// 查询用户是否点赞某个视频
func (dbMgr *manager) QueryUserIsLikeVideo(userId int64, videoId int64) bool {
	var count int64
	if dbMgr.DB.Model(&system.UserLikeVideo{}).Where("user_id = ? AND video_id = ?", userId, videoId).Count(&count); count > 0 {
		return true
	}
	return false
}

// 用户关注
func (dbMgr *manager) UpdateUserWhenFollow(followerId int64, followId int64) error {
	return dbMgr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.UserInfo{}).Where("user_id = ?", followerId).Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
			return err
		}
		if err := tx.Model(&system.UserInfo{}).Where("user_id = ?", followId).Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
			return err
		}
		if err := tx.Create(&system.UserRelation{FollowerId: followerId, FollowId: followId}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 用户取消关注
func (dbMgr *manager) UpdateUserWhenCancelFollow(followerId int64, followId int64) error {
	return dbMgr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.UserInfo{}).Where("user_id = ?", followerId).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.Model(&system.UserInfo{}).Where("user_id = ?", followId).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("follower_id = ? AND follow_id = ?", followerId, followId).Delete(&system.UserRelation{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据userId查询关注的用户
func (dbMgr *manager) QueryFollowUserByUserId(userId int64) ([]*system.UserInfo, error) {
	var followList []*system.UserInfo
	if err := dbMgr.DB.Raw("SELECT * FROM user_info u, user_relation v WHERE v.follower_id = ? AND u.user_id = v.follow_id", userId).Scan(&followList).Error; err != nil {
		return nil, err
	}
	return followList, nil
}

// 查询粉丝
func (dbMgr *manager) QueryFollowerUserByUserId(userId int64) ([]*system.UserInfo, error) {
	var followList []*system.UserInfo
	if err := dbMgr.DB.Raw("SELECT * FROM user_info u, user_relation v WHERE v.follow_id = ? AND u.user_id = v.follower_id", userId).Scan(&followList).Error; err != nil {
		return nil, err
	}
	return followList, nil
}

// 查询好友
func (dbMgr *manager) QueryFriendUserByUserId(userId int64) ([]*system.UserInfo, error) {
	var followIds []int64
	if err := dbMgr.DB.Model(&system.UserRelation{}).Select("follow_id").Where("follower_id = ?", userId).Find(&followIds).Error; err != nil {
		return nil, err
	}
	var friendList []*system.UserInfo
	for _, followId := range followIds {
		if cache.QueryUserIsFollowUser(followId, userId) {
			friend, err := dbMgr.QueryUserByUserId(followId)
			if err != nil {
				return nil, err
			}
			friendList = append(friendList, friend)
		}
	}
	return friendList, nil
}
