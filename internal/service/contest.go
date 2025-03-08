package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"sort"
	"work-project/internal/admin"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Contest interface {
	GetActive(ctx context.Context, userId string) (schema.ContestActivity, error)
	Get(ctx context.Context, request schema.ContestGetRequest) (schema.ContestFullData, error)
	Join(ctx context.Context, data schema.JoinContestRequest) error
	GetDataForSocket(ctx context.Context, userId string) ([]schema.ContestSocketResponse, error)
	Read(ctx context.Context, data schema.ReadContestRequest) (schema.ContestPassBook, error)
	GetPrizes(ctx context.Context, contestId uint) ([]model.ContestPrize, error)
	GetBooks(ctx context.Context, contestId uint) ([]model.ContestBook, error)
	GetBookByID(ctx context.Context, contestBookId uint) (model.ContestBook, error)
	Create(ctx context.Context, data admin.CreateContest) (model.Contest, error)
	Update(ctx context.Context, data admin.UpdateContest) (model.Contest, error)
	Delete(ctx context.Context, contestId uint) error
	GetAll(ctx context.Context) ([]model.Contest, error)
	CreateBook(ctx context.Context, data admin.CreateContestBook) (model.ContestBook, error)
	UpdateBook(ctx context.Context, data admin.UpdateContestBook) (model.ContestBook, error)
	DeleteBook(ctx context.Context, id uint) error
	CreatePrize(ctx context.Context, data admin.CreateContestPrize) (model.ContestPrize, error)
	UpdatePrize(ctx context.Context, data admin.UpdateContestPrize) (model.ContestPrize, error)
	DeletePrize(ctx context.Context, id uint) error
}

type ContestService struct {
	contestRepo            repository.Contest
	contestParticipantRepo repository.ContestParticipant
	contestBookRepo        repository.ContestBook
	contestHistoryRepo     repository.ContestHistory
	contestPrizeRepo       repository.ContestPrize
	storage                repository.StorageClient

	balanceService BalanceService
}

func NewContestService(
	contestRepo repository.Contest,
	contestParticipantRepo repository.ContestParticipant,
	contestBookRepo repository.ContestBook,
	contestHistoryRepo repository.ContestHistory,
	contestPrizeRepo repository.ContestPrize,
	balanceService BalanceService,
) *ContestService {
	return &ContestService{
		contestRepo:            contestRepo,
		contestParticipantRepo: contestParticipantRepo,
		contestBookRepo:        contestBookRepo,
		contestHistoryRepo:     contestHistoryRepo,
		contestPrizeRepo:       contestPrizeRepo,
		balanceService:         balanceService,
	}
}

func (s *ContestService) getResultPrizes(ctx context.Context, userId string) ([]schema.ContestResultPrize, error) {
	contests, err := s.contestRepo.GetEnded(ctx, userId)
	if err != nil {
		return nil, err
	}

	prizes := make([]schema.ContestResultPrize, 0)
	for _, contest := range contests {
		contestParticipant, err := s.contestParticipantRepo.GetByContestAndUserID(ctx, contest.ContestID, userId)
		if err != nil {
			return nil, err
		}
		contestPrize, err := s.contestPrizeRepo.GetByContestIDAndNumber(ctx, contest.ContestID, contestParticipant.Number)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		result := schema.ContestResultPrize{
			ContestId: contest.ContestID,
		}
		if err == nil {
			result.Prize = &contestPrize
		}
		result.Number = contestParticipant.Number
		result.ConsolationPrizeSapphire = contest.ConsolationPrizeSapphire
		prizes = append(prizes, result)
	}

	return prizes, nil
}

func (s *ContestService) Read(ctx context.Context, data schema.ReadContestRequest) (schema.ContestPassBook, error) {
	_, err := s.contestHistoryRepo.GetByContestBookAndUserID(ctx, data.ContestBookId, data.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return schema.ContestPassBook{}, err
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return schema.ContestPassBook{}, fmt.Errorf("user already read this book")
	}

	contestBook, err := s.contestBookRepo.GetByID(ctx, data.ContestBookId)
	if err != nil {
		return schema.ContestPassBook{}, err
	}
	contest, err := s.contestRepo.GetById(ctx, contestBook.ContestID)
	if err != nil {
		return schema.ContestPassBook{}, err
	}
	contestParticipant, err := s.contestParticipantRepo.GetByContestAndUserID(ctx, contest.ContestID, data.UserID)
	if err != nil {
		return schema.ContestPassBook{}, err
	}

	points := int(float64(data.RightQuestions/contestBook.CountOfQuestions) * 10 * 1.5)
	coins := int(float64(data.RightQuestions/contestBook.CountOfQuestions) * 10 * 1.5)
	if contest.CurrentDayNumber() == contestBook.DayNumber {
		coins = int(float64(coins * 4 / 3))
		points *= int(float64(points * 4 / 3))
	}

	_, err = s.contestHistoryRepo.Create(ctx, model.ContestHistory{
		ContestID:     contest.ContestID,
		ContestBookID: contestBook.ContestBookID,
		UserID:        data.UserID,
		Points:        points,
		ReadTime:      data.ReadTime,
	})

	contestParticipant.Points += points
	contestParticipant.ReadTime += data.ReadTime
	contestParticipant, err = s.contestParticipantRepo.Update(ctx, contestParticipant)
	if err != nil {
		return schema.ContestPassBook{}, err
	}

	_, err = s.balanceService.CreateTransaction(ctx, data.UserID, model.Transaction{
		UserId:            data.UserID,
		TransactionType:   string(model.TRANSACTION_TYPE_INCOME),
		Coins:             coins,
		TransactionReason: string(model.TRANSACTION_REASON_CONTEST),
		OuterId:           contestBook.ContestBookID,
	})
	return schema.ContestPassBook{
		Coins:  coins,
		Points: points,
	}, err
}

func (s *ContestService) GetActive(ctx context.Context, userId string) (schema.ContestActivity, error) {
	contests, err := s.contestRepo.GetActive(ctx)
	if err != nil {
		return schema.ContestActivity{}, err
	}

	activeContest := make([]schema.ContestData, 0)
	for _, contest := range contests {
		activeContest = append(activeContest, schema.ContestData{
			ContestID:        contest.ContestID,
			AlreadyJoined:    contest.UserJoined(userId),
			StartDate:        contest.StartTime,
			EndTime:          contest.EndTime,
			CurrentDayNumber: contest.CurrentDayNumber(),
			TotalUsersCount:  len(contest.ContestParticipants),
		})
	}

	endedContestPrizes, err := s.getResultPrizes(ctx, userId)
	if err != nil {
		return schema.ContestActivity{}, err
	}

	return schema.ContestActivity{
		Active: activeContest,
		Ended:  endedContestPrizes,
	}, nil
}

func (s *ContestService) Get(ctx context.Context, request schema.ContestGetRequest) (schema.ContestFullData, error) {
	contest, err := s.contestRepo.GetById(ctx, request.ContestID)
	if err != nil {
		return schema.ContestFullData{}, err
	}
	result := schema.ContestFullData{
		ContestData: schema.ContestData{
			ContestID:        contest.ContestID,
			AlreadyJoined:    contest.UserJoined(request.UserId),
			StartDate:        contest.StartTime,
			EndTime:          contest.EndTime,
			CurrentDayNumber: contest.CurrentDayNumber(),
			TotalUsersCount:  len(contest.ContestParticipants),
		},
		Books: contest.ContestBooks.GetContestBookSchema(request.UserId),
	}

	return result, nil
}

func (s *ContestService) Join(ctx context.Context, data schema.JoinContestRequest) error {
	contest, err := s.contestRepo.GetById(ctx, data.ContestId)
	if err != nil {
		return err
	}
	if contest.UserJoined(data.UserID) {
		return fmt.Errorf("user already joined")
	}

	_, err = s.contestParticipantRepo.Create(ctx, model.ContestParticipant{
		ContestID: contest.ContestID,
		UserID:    data.UserID,
		ReadTime:  0,
		Points:    0,
	})
	return err
}

func (s *ContestService) GetDataForSocket(ctx context.Context, userId string) ([]schema.ContestSocketResponse, error) {
	contests, err := s.contestRepo.GetActiveJoinedByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	var result []schema.ContestSocketResponse
	for _, contest := range contests {
		participants := contest.ContestParticipants

		sort.Slice(participants, func(i, j int) bool {
			if participants[i].Points == participants[j].Points {
				return participants[i].ReadTime < participants[j].ReadTime
			}
			return participants[i].Points > participants[j].Points
		})

		var topUsers []schema.ContestUserSocketData
		var currentUserData schema.ContestUserSocketData

		for index, participant := range participants {
			userData := schema.ContestUserSocketData{
				UserId:        participant.UserID,
				Nickname:      participant.User.Profile.UserName,
				Number:        index + 1,
				Points:        participant.Points,
				TotalReadTime: participant.ReadTime,
			}

			if index < 5 {
				topUsers = append(topUsers, userData)
			}

			if participant.UserID == userId {
				currentUserData = userData
			}
		}

		result = append(result, schema.ContestSocketResponse{
			ContestId:   contest.ContestID,
			EndTime:     contest.EndTime,
			StartTime:   contest.StartTime,
			TopUsers:    topUsers,
			CurrentUser: currentUserData,
		})
	}

	return result, nil
}

func (s *ContestService) GetPrizes(ctx context.Context, contestId uint) ([]model.ContestPrize, error) {
	return s.contestPrizeRepo.GetByContestID(ctx, contestId)
}

func (s *ContestService) GetBooks(ctx context.Context, contestId uint) ([]model.ContestBook, error) {
	return s.contestBookRepo.GetByContestID(ctx, contestId)

}
func (s *ContestService) GetBookByID(ctx context.Context, contestBookId uint) (model.ContestBook, error) {
	return s.contestBookRepo.GetByID(ctx, contestBookId)
}
func (s *ContestService) Create(ctx context.Context, data admin.CreateContest) (model.Contest, error) {
	return s.contestRepo.Create(ctx, model.Contest{
		StartTime:                data.StartTime,
		EndTime:                  data.EndTime,
		IsActive:                 data.IsActive,
		ConsolationPrizeSapphire: data.ConsolationPrizeSapphire,
	})
}
func (s *ContestService) Update(ctx context.Context, data admin.UpdateContest) (model.Contest, error) {
	contest, err := s.contestRepo.GetById(ctx, data.ContestID)
	if err != nil {
		return model.Contest{}, err
	}
	if data.EndTime != nil {
		contest.EndTime = *data.EndTime
	}
	if data.StartTime != nil {
		contest.EndTime = *data.StartTime
	}
	if data.IsActive != nil {
		contest.IsActive = *data.IsActive
	}
	if data.ConsolationPrizeSapphire != nil {
		contest.ConsolationPrizeSapphire = *data.ConsolationPrizeSapphire
	}

	return s.contestRepo.Update(ctx, contest)
}

func (s *ContestService) Delete(ctx context.Context, contestId uint) error {
	return s.contestRepo.Delete(ctx, contestId)
}

func (s *ContestService) GetAll(ctx context.Context) ([]model.Contest, error) {
	return s.contestRepo.GetAll(ctx)
}

func (s *ContestService) CreateBook(ctx context.Context, data admin.CreateContestBook) (model.ContestBook, error) {
	contestBook := model.ContestBook{
		ContestID:        data.ContestID,
		DayNumber:        data.DayNumber,
		Title:            data.Title,
		TitleKz:          data.TitleKz,
		TitleEn:          data.TitleEn,
		Description:      data.Description,
		DescriptionKz:    data.DescriptionKz,
		DescriptionEn:    data.DescriptionEn,
		Body:             data.Body,
		BodyEn:           data.BodyEn,
		BodyKz:           data.BodyKz,
		Point:            data.Point,
		ContestCoins:     data.ContestCoins,
		CountOfQuestions: data.CountOfQuestions,
	}

	if data.Image != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_CONTEST), data.Image.FileName, data.Image.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "contest book name", contestBook.Title)
			return model.ContestBook{}, err
		}
		contestBook.PhotoPath = &file
	}

	return s.contestBookRepo.Create(ctx, contestBook)
}

func (s *ContestService) UpdateBook(ctx context.Context, data admin.UpdateContestBook) (model.ContestBook, error) {
	contestBook, err := s.contestBookRepo.GetByID(ctx, data.ContestBookID)
	if err != nil {
		return model.ContestBook{}, err
	}

	if data.DayNumber != nil {
		contestBook.DayNumber = *data.DayNumber
	}
	if data.TitleKz != nil {
		contestBook.TitleKz = *data.TitleKz
	}
	if data.Title != nil {
		contestBook.Title = *data.Title
	}
	if data.TitleEn != nil {
		contestBook.TitleEn = *data.TitleEn
	}
	if data.Description != nil {
		contestBook.Description = *data.Description
	}
	if data.DescriptionEn != nil {
		contestBook.DescriptionEn = *data.DescriptionEn
	}
	if data.DescriptionKz != nil {
		contestBook.DescriptionKz = *data.DescriptionKz
	}
	if data.Body != nil {
		contestBook.Body = *data.Body
	}
	if data.BodyEn != nil {
		contestBook.BodyEn = *data.BodyEn
	}
	if data.BodyKz != nil {
		contestBook.BodyKz = *data.BodyKz
	}

	if data.Point != nil {
		contestBook.Point = *data.Point
	}
	if data.ContestCoins != nil {
		contestBook.ContestCoins = *data.ContestCoins
	}
	if data.Point != nil {
		contestBook.CountOfQuestions = *data.CountOfQuestions
	}

	if data.Image != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_CONTEST), data.Image.FileName, data.Image.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "contest book name", contestBook.Title)
			return model.ContestBook{}, err
		}
		contestBook.PhotoPath = &file
	}

	return s.contestBookRepo.Update(ctx, contestBook)
}

func (s *ContestService) DeleteBook(ctx context.Context, id uint) error {
	return s.contestBookRepo.Delete(ctx, id)
}

func (s *ContestService) CreatePrize(ctx context.Context, data admin.CreateContestPrize) (model.ContestPrize, error) {
	contestPrize := model.ContestPrize{
		ContestID: data.ContestID,
		Number:    data.Number,
		PrizeName: data.PrizeName,
	}

	if data.Image != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_CONTEST), data.Image.FileName, data.Image.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "contest prize name", contestPrize.PrizeName)
			return model.ContestPrize{}, err
		}
		contestPrize.PhotoPath = &file
	}

	return s.contestPrizeRepo.Create(ctx, contestPrize)
}

func (s *ContestService) UpdatePrize(ctx context.Context, data admin.UpdateContestPrize) (model.ContestPrize, error) {
	contestPrize, err := s.contestPrizeRepo.GetByID(ctx, data.ContestPrizeID)
	if err != nil {
		return model.ContestPrize{}, err
	}

	if data.PrizeName != nil {
		contestPrize.PrizeName = *data.PrizeName
	}
	if data.Number != nil {
		contestPrize.Number = *data.Number
	}

	if data.Image != nil {
		file, err := s.storage.CreateImageFromBase64(ctx, string(model.BUCKET_NAME_CONTEST), data.Image.FileName, data.Image.File)
		if err != nil {
			log.Println(ctx, "some err while create image", "err", err, "contest prize name", contestPrize.PrizeName)
			return model.ContestPrize{}, err
		}
		contestPrize.PhotoPath = &file
	}

	return s.contestPrizeRepo.Update(ctx, contestPrize)
}
func (s *ContestService) DeletePrize(ctx context.Context, id uint) error {
	return s.contestPrizeRepo.Delete(ctx, id)
}
