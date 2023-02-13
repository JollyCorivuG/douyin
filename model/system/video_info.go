package system

import "time"

// 视频信息
type VideoInfo struct {
	VideoId       int64     `json:"id" gorm:"primarykey;autoincrement:false"`
	AuthorId      int64     `json:"-" gorm:"autoincrement:false"`
	VideoAuthor   *UserInfo `json:"author" gorm:"foreignkey:AuthorId"`
	PlayUrl       string    `json:"play_url" gorm:"not null"`
	CoverUrl      string    `json:"cover_url" gorm:"not null"`
	FavoriteCount int       `json:"favorite_count" gorm:"autoincrement:false"`
	CommentCount  int       `json:"comment_count" gorm:"autoincrement:false"`
	IsFavorite    bool      `json:"is_favorite" gorm:"-"`
	Title         string    `json:"title" gorm:"not null"`
	CreateAt      time.Time `json:"-" gorm:"not null"`
}
