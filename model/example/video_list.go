package example

import "douyin/model/system"

// feed返回的videoList
type VideoList struct {
	NextTime int64              `json:"next_time"` //本次返回的视频中, 发布最早的时间, 作为下次请求时的latest_time
	Videos   []system.VideoInfo `json:"video_list"`
}
