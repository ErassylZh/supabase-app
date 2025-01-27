package service

import (
	"context"
	"fmt"
	"sort"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/schema"
)

type Contest interface {
	GetActive(ctx context.Context, userId string) ([]schema.ContestData, error)
	Get(ctx context.Context, request schema.ContestGetRequest) (schema.ContestFullData, error)
	Join(ctx context.Context, data schema.JoinContestRequest) error
	GetDataForSocket(ctx context.Context, userId string) ([]schema.ContestSocketResponse, error)
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
			TotalUsersCount:  len(contest.ContestParticipants),
		})
	}
	return result, nil
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
