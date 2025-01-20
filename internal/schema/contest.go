package schema

import "time"

type ContestData struct {
	ContestID        uint      `json:"contest_id"`
	AlreadyJoined    bool      `json:"already_joined"`
	StartDate        time.Time `json:"start_date"`
	EndTime          time.Time `json:"end_time"`
	CurrentDayNumber int       `json:"current_day_number"`
}

type JoinContestRequest struct {
	ContestId uint   `json:"contest_id"`
	UserID    string `json:"-"`
}
