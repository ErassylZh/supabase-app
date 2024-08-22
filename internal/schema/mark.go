package schema

type MarkCreateRequest struct {
	MarkID string `json:"-"`
	Token  string `json:"token"`
}
