package schema

type CreateMark struct {
	PostId uint   `json:"post_id"`
	UserID string `json:"-"`
}
