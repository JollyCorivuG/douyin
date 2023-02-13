package system

import "gorm.io/gorm"

// 存储userid -> videoId
type UserLikeVideo struct {
	gorm.Model
	UserId  int64 `gorm:"autoincrement:false"`
	VideoId int64 `gorm:"autoincrement:false"`
}
