package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ContestParticipant interface {
	GetByContestID(ctx context.Context, contestId uint) ([]model.ContestParticipant, error)
	Create(ctx context.Context, contestParticipant model.ContestParticipant) (model.ContestParticipant, error)
	Update(ctx context.Context, contestParticipant model.ContestParticipant) (model.ContestParticipant, error)
}

type ContestParticipantDB struct {
	db *gorm.DB
}

func NewContestParticipantDB(db *gorm.DB) *ContestParticipantDB {
	return &ContestParticipantDB{db: db}
}

func (r *ContestParticipantDB) GetByContestID(ctx context.Context, contestId uint) (contestParticipants []model.ContestParticipant, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestParticipant{})
	err = q.Where("contest_id = ?", contestId).
		Find(&contestParticipants).
		Error
	if err != nil {
		return contestParticipants, err
	}
	return contestParticipants, nil
}

func (r *ContestParticipantDB) Create(ctx context.Context, contestParticipant model.ContestParticipant) (model.ContestParticipant, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contestParticipant).
		Error
	if err != nil {
		return contestParticipant, err
	}
	return contestParticipant, nil
}

func (r *ContestParticipantDB) Update(ctx context.Context, contestParticipant model.ContestParticipant) (model.ContestParticipant, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Where("contest_participant_id = ?", contestParticipant.ContestParticipantID).
		Save(&contestParticipant).
		Error
	if err != nil {
		return contestParticipant, err
	}
	return contestParticipant, nil
}
