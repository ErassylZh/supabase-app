package schema

import "work-project/internal/model"

type PostResponse struct {
	model.Post
	IsMarked bool `json:"is_marked"`
}
