package service

import (
	"context"
	"fmt"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Contest interface {
	GetActive(ctx context.Context, userId string) ([]schema.ContestData, error)
	Join(ctx context.Context, data schema.JoinContestRequest) error
}

type ContestService struct {
	contestRepo            repository.Contest
	contestParticipantRepo repository.ContestParticipant
	contestBookRepo        repository.ContestBook
}

func NewContestService(contestRepo repository.Contest, contestParticipantRepo repository.ContestParticipant, contestBookRepo repository.ContestBook) *ContestService {
	return &ContestService{contestRepo: contestRepo, contestParticipantRepo: contestParticipantRepo, contestBookRepo: contestBookRepo}
}

func (s *ContestService) GetActive(ctx context.Context, userId string) ([]schema.ContestData, error) {
	contests, err := s.contestRepo.GetActive(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]schema.ContestData, 0)
	for _, contest := range contests {
		result = append(result, schema.ContestData{
			ContestID:        contest.ContestID,
			AlreadyJoined:    contest.UserJoined(userId),
			StartDate:        contest.StartTime,
			EndTime:          contest.EndTime,
			CurrentDayNumber: contest.CurrentDayNumber(),
		})
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
