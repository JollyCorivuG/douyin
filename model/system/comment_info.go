package system

type CommentInfo struct {
	CommentId     int64     `json:"id" gorm:"primarykey;autoincrement:false"`
	PosterId      int64     `json:"-" gorm:"autoincrement:false"`
	CommentPoster *UserInfo `json:"user" gorm:"foreignkey:PosterId"`
	VideoId       int64     `json:"-" gorm:"autoincrement:false"`
	Content       string    `json:"content" gorm:"type:text"`
	CreateDate    string    `json:"create_date"`
}
