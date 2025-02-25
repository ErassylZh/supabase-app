package schema

type UserUpdate struct {
	UserID      string  `json:"-"`
	Nickname    string  `json:"nickname"`
	AvatarPhoto *string `json:"avatar_photo"`
}
