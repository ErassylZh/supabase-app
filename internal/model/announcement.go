package model

import "time"

type Announcement struct {
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	CreatedAt      time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	AnnouncementID uint      `gorm:"primaryKey;column:announcement_id" json:"id"`
	Text           string    `gorm:"column:text" json:"text"`
	Header         string    `gorm:"column:header" json:"header"`
	AuthorID       uint      `gorm:"column:author_id" json:"author_id"`
	Date           string    `gorm:"-" json:"date"`

	Author User `json:"author"`
}

func (o *Announcement) TableName() string {
	return "student_repo.announcement"
}
