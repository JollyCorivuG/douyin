package system

import "gorm.io/gorm"

type UserRelation struct {
	gorm.Model
	FollowerId int64 `gorm:"autoincrement:false"`
	FollowId   int64 `gorm:"autoincrement:false"`
}
