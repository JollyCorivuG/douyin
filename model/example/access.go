package example

// 用户权限, 含userId和token
type Access struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
