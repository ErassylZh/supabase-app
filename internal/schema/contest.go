package schema

import (
	"time"
	"work-project/internal/model"
)

type ContestData struct {
	ContestID        uint      `json:"contest_id"`
	AlreadyJoined    bool      `json:"already_joined"`
	StartDate        time.Time `json:"start_date"`
	EndTime          time.Time `json:"end_time"`
	CurrentDayNumber int       `json:"current_day_number"`
	TotalUsersCount  int       `json:"total_users_count"`
}

type ContestGetRequest struct {
	ContestID uint   `form:"contest_id"`
	UserId    string `form:"-" json:"-"`
}

type ContestFullData struct {
	ContestData
	Books []ContestBookData `json:"books"`
}

type ContestBookData struct {
	model.ContestBook
	Readed bool
}

type JoinContestRequest struct {
	ContestId uint   `json:"contest_id"`
	UserID    string `json:"-"`
}

type ContestSocketResponse struct {
	ContestId   uint                    `json:"contest_id"`
	EndTime     time.Time               `json:"end_time"`
	StartTime   time.Time               `json:"start_time"`
	TopUsers    []ContestUserSocketData `json:"top_users"`
	CurrentUser ContestUserSocketData   `json:"current_user"`
}

type ContestUserSocketData struct {
	UserId        string `json:"user_id"`
	Nickname      string `json:"nickname"`
	Number        int    `json:"number"`
	Points        int    `json:"points"`
	TotalReadTime int    `json:"total_read_time"`
}
