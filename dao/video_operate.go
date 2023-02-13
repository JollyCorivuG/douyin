package dao

import (
	"douyin/model/system"
	"time"

	"gorm.io/gorm"
)

// 添加视频到数据库
func (dbMgr *manager) AddVideo(videoInfo *system.VideoInfo) error {
	return dbMgr.DB.Create(videoInfo).Error
}

// 根据用户限制最新的投稿时间戳返回视频
func (dbMgr *manager) QueryVideosByLimit(limitNum int, latestTime time.Time) ([]system.VideoInfo, error) {
	var videoInfos []system.VideoInfo
	if err := dbMgr.DB.Model(&system.VideoInfo{}).
		Order("create_at DESC").
		Where("create_at < ?", latestTime).
		Limit(limitNum).
		Find(&videoInfos).Error; err != nil {
		return nil, err
	}
	return videoInfos, nil
}

// 根据用户id查询此用户发布的视频
func (dbMgr *manager) QueryVideosByUserId(userId int64) ([]*system.VideoInfo, error) {
	var videoInfos []*system.VideoInfo
	if err := dbMgr.DB.Where("author_id = ?", userId).Find(&videoInfos).Error; err != nil {
		return nil, err
	}
	return videoInfos, nil
}

// 用户点赞视频
func (dbMgr *manager) UpdateVideoWhenLike(userId int64, videoId int64) error {
	return dbMgr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.VideoInfo{}).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
			return err
		}
		if err := tx.Create(&system.UserLikeVideo{UserId: userId, VideoId: videoId}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 用户取消点赞视频
func (dbMgr *manager) UpdateVideoWhenCancelLike(userId int64, videoId int64) error {
	return dbMgr.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&system.VideoInfo{}).Where("video_id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&system.UserLikeVideo{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据userId查询点赞的视频
func (dbMgr *manager) QueryLikeVideosByUserId(userId int64) ([]*system.VideoInfo, error) {
	// 多表查询, 将user_like_video和video_info两张表联合起来查询
	var videoList []*system.VideoInfo
	if err := dbMgr.DB.Raw("SELECT u.* FROM video_info u, user_like_video v WHERE v.user_id = ? AND u.video_id = v.video_id", userId).Scan(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}
