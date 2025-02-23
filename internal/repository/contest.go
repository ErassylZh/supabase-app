package repository

import (
	"context"
	"gorm.io/gorm"
	"time"
	"work-project/internal/model"
)

type Contest interface {
	GetActive(ctx context.Context) ([]model.Contest, error)
	GetById(ctx context.Context, contestId uint) (model.Contest, error)
	Create(ctx context.Context, contest model.Contest) (model.Contest, error)
	Update(ctx context.Context, contest model.Contest) (model.Contest, error)
	GetActiveJoinedByUser(ctx context.Context, id string) ([]model.Contest, error)
	GetEnded(ctx context.Context, userId string) (contests []model.Contest, err error)
	GetAll(ctx context.Context) ([]model.Contest, error)
	CreateMany(ctx context.Context, contests []model.Contest) ([]model.Contest, error)
	UpdateMany(ctx context.Context, contests []model.Contest) ([]model.Contest, error)
}

type ContestDB struct {
	db *gorm.DB
}

func NewContestDB(db *gorm.DB) *ContestDB {
	return &ContestDB{db: db}
}

func (r *ContestDB) GetAll(ctx context.Context) (contests []model.Contest, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Contest{})
	err = q.Find(&contests).
		Error
	if err != nil {
		return contests, err
	}
	return contests, nil
}

func (r *ContestDB) GetActive(ctx context.Context) (contests []model.Contest, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Contest{})
	err = q.Where("is_active and end_time > ?", time.Now()).
		Preload("ContestParticipants").
		Find(&contests).
		Error
	if err != nil {
		return contests, err
	}
	return contests, nil
}

func (r *ContestDB) GetEnded(ctx context.Context, userId string) (contests []model.Contest, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Contest{})
	err = q.Joins("JOIN public.contest_participant on public.contest_participant.contest_id = public.contest.contest_id").
		Where("public.contest_participant.user_id = ?", userId).
		Where("not public.contest.is_active and public.contest.end_time < ?", time.Now()).
		Find(&contests).
		Error
	if err != nil {
		return contests, err
	}
	return contests, nil
}

func (r *ContestDB) GetById(ctx context.Context, contestId uint) (contest model.Contest, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Contest{})
	err = q.Where("contest_id ?", contestId).
		Preload("ContestParticipants").
		Preload("ContestBooks.ContestHistory").
		First(&contest).
		Error
	if err != nil {
		return contest, err
	}
	return contest, nil
}

func (r *ContestDB) Create(ctx context.Context, contest model.Contest) (model.Contest, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contest).
		Error
	if err != nil {
		return contest, err
	}
	return contest, nil
}

func (r *ContestDB) Update(ctx context.Context, contest model.Contest) (model.Contest, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Where("contest_id = ?", contest.ContestID).
		Save(&contest).
		Error
	if err != nil {
		return contest, err
	}
	return contest, nil
}

func (r *ContestDB) GetActiveJoinedByUser(ctx context.Context, userId string) (contests []model.Contest, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Contest{})
	err = q.Joins("inner join public.contest_participant on public.contest.contest_id =public.contest_participant.contest_id ").
		Where("public.contest_participant.user_id = ?", userId).
		Where("public.contest.is_active and public.contest.end_time > ? ", time.Now()).
		Preload("ContestParticipants").
		Preload("ContestParticipants.User.Profile").
		Find(&contests).
		Error
	if err != nil {
		return contests, err
	}
	return contests, nil
}

func (r *ContestDB) CreateMany(ctx context.Context, contests []model.Contest) ([]model.Contest, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.Collection{})
	err := q.Create(&contests).
		Error
	if err != nil {
		return contests, err
	}
	return contests, nil
}

func (r *ContestDB) UpdateMany(ctx context.Context, contests []model.Contest) ([]model.Contest, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, contest := range contests {
			if err := tx.Model(&model.Contest{}).Where("contest_id = ?", contest.ContestID).Updates(map[string]interface{}{
				"is_active":                  contest.IsActive,
				"start_time":                 contest.StartTime,
				"end_time":                   contest.EndTime,
				"code":                       contest.Code,
				"consolation_prize_sapphire": contest.ConsolationPrizeSapphire,
			}).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return contests, nil
}
