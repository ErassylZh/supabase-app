package repository

import (
	"context"
	"gorm.io/gorm"
)

type BadWord interface {
	GetAll(ctx context.Context) ([]string, error)
}
type BadWordDB struct {
	db *gorm.DB
}

func NewBadWordDB(db *gorm.DB) *BadWordDB {
	return &BadWordDB{db: db}
}

func (r *BadWordDB) GetAll(ctx context.Context) ([]string, error) {
	var words []string
	err := r.db.
		Table("public.bad_words").
		Pluck("word", &words).
		Error
	if err != nil {
		return nil, err
	}
	return words, nil
}
