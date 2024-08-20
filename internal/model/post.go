package model

import (
	"fmt"
	"time"
)

type PostLanguage string

const (
	POST_LANGUAGE_KZ PostLanguage = "kz"
	POST_LANGUAGE_EN PostLanguage = "en"
	POST_LANGUAGE_RU PostLanguage = "ru"
)

type PostStatus string

const (
	POST_STATUS_PUBLISH PostStatus = "Publish"
	POST_STATUS_DRAFT   PostStatus = "Draft"
)

type PostRatingStatus string

const (
	POST_RATING_STATUS_POPULAR PostRatingStatus = "Popular"
	POST_RATING_STATUS_NORMAL  PostRatingStatus = "Normal"
)

type PostImageType string

const (
	POST_IMAGE_TYPE_IMAGE PostImageType = "image"
	POST_IMAGE_TYPE_LOGO  PostImageType = "logo"
)

type Post struct {
	CreatedAt    time.Time
	PostID       uint    `gorm:"primaryKey;column:post_id" json:"post_id"`
	Company      *string `gorm:"column:company" json:"company"`
	Logo         *string `gorm:"column:logo" json:"logo"`
	Language     *string `gorm:"column:language" json:"language"`
	Title        string  `gorm:"column:title" json:"title"`
	Description  *string `gorm:"column:description" json:"description"`
	Status       *string `gorm:"column:status" json:"status"`
	Image        *string `gorm:"column:image" json:"image"`
	Hashtags     *string `gorm:"column:hashtags" json:"hashtags"`
	HashtagName  *string `gorm:"column:hashtag_name" json:"hashtag_name"`
	Body         *string `gorm:"column:body" json:"body"`
	ReadTime     *int    `gorm:"column:read_time" json:"read_time"`
	Point        *int    `gorm:"column:point" json:"point"`
	QuizTime     *int    `gorm:"column:quiz_time" json:"quiz_time"`
	RatingStatus *string `gorm:"column:rating_status" json:"rating_status"`
	Uuid         string  `gorm:"column:uuid" json:"uuid"`
}

func (p Post) TableName() string {
	return fmt.Sprintf("public.post")
}
