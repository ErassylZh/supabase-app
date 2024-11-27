package model

import "time"

type Auction struct {
	CreatedAt           time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at" json:"updated_at"`
	AuctionID           uint      `gorm:"column:auction_id" json:"auction_id"`
	StartTime           time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime             time.Time `gorm:"column:end_time" json:"end_time"`
	EntryPrice          int       `gorm:"column:entry_price" json:"entry_price"`
	PostId              *uint     `gorm:"column:post_id" json:"post_id"`
	MaxParticipantCount int       `gorm:"column:max_participant_count" json:"max_participant_count"`
	ProductId           *uint     `gorm:"column:product_id" json:"product_id"`
}
