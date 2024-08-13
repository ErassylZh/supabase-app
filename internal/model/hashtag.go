package model

import "time"

type Hashtag struct {
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	HashtagId         uint      `gorm:"column:hashtag_id" json:"hashtag_id"`
	AirtableHashtagId *string   `gorm:"column:airtable_hashtag_id" json:"airtable_hashtag_id"`
	Name              string    `gorm:"column:name" json:"name"`
}
