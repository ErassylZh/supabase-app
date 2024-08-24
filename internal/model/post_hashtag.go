package model

import "fmt"

type PostHashtag struct {
	PostId    uint `gorm:"column:post_id" json:"post_id"`
	HashtagId uint `gorm:"column:hashtag_id" json:"hashtag_id"`
}

func (p PostHashtag) TableName() string {
	return fmt.Sprintf("public.post_hashtag")
}
