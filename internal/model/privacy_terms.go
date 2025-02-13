package model

import (
	"fmt"
)

type PrivacyTerms struct {
	PrivacyTermsID uint   `gorm:"primaryKey;column:privacy_terms_id" json:"privacy_terms_id"`
	Type           string `gorm:"column:type" json:"type"`
	Content        string `gorm:"column:content" json:"content"`
	Status         string `gorm:"column:status" json:"status"`
}

func (m PrivacyTerms) TableName() string {
	return fmt.Sprintf("public.privacy_terms")
}
