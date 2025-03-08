package repository

import (
	"context"
	"gorm.io/gorm"
	"work-project/internal/model"
)

type ContestBook interface {
	GetByContestID(ctx context.Context, contestId uint) ([]model.ContestBook, error)
	GetByID(ctx context.Context, contestBookId uint) (model.ContestBook, error)
	CreateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error)
	UpdateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error)
	GetAll(ctx context.Context) ([]model.ContestBook, error)
	Create(ctx context.Context, book model.ContestBook) (model.ContestBook, error)
	Update(ctx context.Context, book model.ContestBook) (model.ContestBook, error)
	Delete(ctx context.Context, id uint) error
}

type ContestBookDB struct {
	db *gorm.DB
}

func NewContestBookDB(db *gorm.DB) *ContestBookDB {
	return &ContestBookDB{db: db}
}

func (r *ContestBookDB) GetAll(ctx context.Context) (contestBooks []model.ContestBook, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err = q.Find(&contestBooks).
		Error
	if err != nil {
		return contestBooks, err
	}
	return contestBooks, nil
}

func (r *ContestBookDB) GetByID(ctx context.Context, contestBookId uint) (contestBook model.ContestBook, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err = q.Where("contest_book_id = ?", contestBookId).
		First(&contestBook).
		Error
	if err != nil {
		return contestBook, err
	}
	return contestBook, nil
}

func (r *ContestBookDB) GetByContestID(ctx context.Context, contestId uint) (contestBooks []model.ContestBook, err error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err = q.Where("contest_id = ?", contestId).
		Find(&contestBooks).
		Error
	if err != nil {
		return contestBooks, err
	}
	return contestBooks, nil
}

func (r *ContestBookDB) CreateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err := q.Create(&contestBooks).
		Error
	if err != nil {
		return contestBooks, err
	}
	return contestBooks, nil
}

func (r *ContestBookDB) UpdateMany(ctx context.Context, contestBooks []model.ContestBook) ([]model.ContestBook, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, contestBook := range contestBooks {
			if err := tx.Model(&model.ContestBook{}).Where("contest_book_id = ?", contestBook.ContestBookID).Updates(&contestBook).Error; err != nil {
				return err // Если ошибка — транзакция отменяется
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return contestBooks, nil
}

func (r *ContestBookDB) Create(ctx context.Context, book model.ContestBook) (model.ContestBook, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err := q.Create(&book).
		Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *ContestBookDB) Update(ctx context.Context, book model.ContestBook) (model.ContestBook, error) {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err := q.Where("contest_book_id = ?", book.ContestBookID).
		Save(&book).
		Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *ContestBookDB) Delete(ctx context.Context, id uint) error {
	db := r.db.WithContext(ctx)
	q := db.Model(&model.ContestBook{})
	err := q.Where("contest_book_id = ?", id).
		Delete(&model.ContestBook{}).
		Error
	if err != nil {
		return err
	}
	return nil
}
