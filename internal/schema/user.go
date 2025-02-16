package schema

type UserUpdate struct {
	UserID   string `json:"-"`
	Nickname string `json:"nickname"`
}
