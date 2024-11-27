package schema

import "work-project/internal/model"

type PostResponse struct {
	model.Post
	IsMarked      bool  `json:"is_marked"`
	MarkId        *uint `json:"mark_id"`
	IsAlreadyRead bool  `json:"is_already_read"`
	QuizPassed    bool  `json:"quiz_passed"`
}

type PostResponseByGroup struct {
	Bestsellers []PostResponse `json:"bestsellers"`
	Partners    []PostResponse `json:"partners"`
	Other       []PostResponse `json:"-"`
}

type ReadPost struct {
	PostId uint `json:"post_id"`
}

type PassQuizPost struct {
	PostId        uint `json:"post_id"`
	QuizPoints    *int `json:"quiz_points"`
	QuizSapphires *int `json:"quiz_sapphires"`
}

type GetListingFilter struct {
	HashtagIds    []uint
	CollectionIds []uint
	Search        *string
	Language      *string
	PostIds       []uint
}
