package schema

type CreateMark struct {
	PostId uint   `json:"post_id"`
	UserID string `json:"-"`
}

type MarkResponse struct {
	MarkID uint         `json:"mark_id"`
	UserID string       `json:"user_id"`
	PostID uint         `json:"post_id"`
	Post   PostResponse `json:"post"`
}
