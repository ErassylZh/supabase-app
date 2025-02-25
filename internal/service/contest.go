package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sort"
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
}

type ContestService struct {
	contestRepo            repository.Contest
	contestParticipantRepo repository.ContestParticipant
	contestBookRepo        repository.ContestBook
	contestHistoryRepo     repository.ContestHistory
	contestPrizeRepo       repository.ContestPrize

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
