package system

type ChatMessage struct {
	MessageId  int64  `json:"id" gorm:"primarykey;autoincrement:false"`
	ToUserId   int64  `json:"to_user_id" gorm:"autoincrement:false"`
	FromUserId int64  `json:"from_user_id" gorm:"autoincrement:false"`
	Content    string `json:"content" gorm:"type:text"`
	CreateTime int64 `json:"create_time" gorm:"autoincrement:false"`
}
