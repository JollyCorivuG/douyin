package system

// 系统用户信息
type UserInfo struct {
	UserId        int64  `json:"id" gorm:"primarykey;autoincrement:false"`
	UserName      string `json:"name" gorm:"type:varchar(32);not null;unique"`
	PassWord      string `json:"-" gorm:"not null"`
	FollowCount   int    `json:"follow_count" gorm:"autoincrement:false"`
	FollowerCount int    `json:"follower_count" gorm:"autoincrement:false"`
	WorkCount     int    `json:"work_count" gorm:"autoincrement:false"`
	FavoriteCount int    `json:"favorite_count" gorm:"autoincrement:false"`
	IsFollow      bool   `json:"is_follow" gorm:"-"`
	// PublishVideos []*VideoInfo `json:"-" gorm:"foreignkey:VideoId"`
}
