package datatypes

type UserFormated struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"is_admin"`
}

const (
	UserStatusNormal int64 = iota + 1
	UserStatusClosed
)
