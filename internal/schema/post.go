package schema

import "work-project/internal/model"

type PostResponse struct {
	model.Post
	IsMarked      bool  `json:"is_marked"`
	MarkId        *uint `json:"mark_id"`
	IsAlreadyRead bool  `json:"is_already_read"`
}

type PostResponseByGroup struct {
	Bestsellers []PostResponse `json:"bestsellers"`
	Partners    []PostResponse `json:"partners"`
	Other       []PostResponse `json:"-"`
}
