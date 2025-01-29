package schema

import "work-project/internal/model"

type PostResponse struct {
	model.Post
	IsMarked      bool   `json:"is_marked"`
	MarkId        *uint  `json:"mark_id"`
	IsAlreadyRead bool   `json:"is_already_read"`
	QuizPassed    bool   `json:"quiz_passed"`
	PostType      string `json:"post_type"`
}

type PostResponseByGroup struct {
	Bestsellers     []PostResponse `json:"bestsellers"`
	Partners        []PostResponse `json:"partners"`
	ContinueReading []PostResponse `json:"continue_reading"`
	Other           []PostResponse `json:"-"`
}

type ReadPost struct {
	PostId uint `json:"post_id"`
	End    bool `json:"end"`
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
	Pagination
}

type ArchivePost struct {
	model.Post
	EarnedCoins     *int `json:"earned_coins"`
	EarnedSapphires *int `json:"earned_sapphires"`
	QuizPassed      bool `json:"quiz_passed"`
}

type ReadPostRequest struct {
	PostId     uint   `json:"post_id"`
	EndReading bool   `json:"end_reading"`
	UserId     string `json:"-"`
}

type PostGetRequest struct {
	HashtagId    uint   `form:"hashtag_id"`
	CollectionId uint   `form:"collection_id"`
	Language     string `form:"language"`
	Pagination
}
