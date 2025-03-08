package admin

import "time"

type CreateContest struct {
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	IsActive                 bool      `json:"is_active"`
	ConsolationPrizeSapphire int       `json:"consolation_prize_sapphire"`
}

type UpdateContest struct {
	ContestID                uint       `json:"contest_id"`
	StartTime                *time.Time `json:"start_time"`
	EndTime                  *time.Time `json:"end_time"`
	IsActive                 *bool      `json:"is_active"`
	ConsolationPrizeSapphire *int       `json:"consolation_prize_sapphire"`
}

type CreateContestBook struct {
	ContestID        uint   `json:"contest_id"`
	DayNumber        int    `json:"day_number"`
	ContestCoins     int    `json:"contest_coins"`
	Point            int    `json:"point"`
	Title            string `json:"title"`
	TitleKz          string `json:"title_kz"`
	TitleEn          string `json:"title_en"`
	Description      string `json:"description"`
	DescriptionKz    string `json:"description_kz"`
	DescriptionEn    string `json:"description_en"`
	Status           string `json:"status"`
	Body             string `json:"body"`
	BodyKz           string `json:"body_kz"`
	BodyEn           string `json:"body_en"`
	CountOfQuestions int    `json:"count_of_questions"`

	Image *Image `json:"image"`
}

type UpdateContestBook struct {
	ContestBookID    uint    `json:"contest_book_id"`
	DayNumber        *int    `json:"day_number"`
	ContestCoins     *int    `json:"contest_coins"`
	Point            *int    `json:"point"`
	Title            *string `json:"title"`
	TitleKz          *string `json:"title_kz"`
	TitleEn          *string `json:"title_en"`
	Description      *string `json:"description"`
	DescriptionKz    *string `json:"description_kz"`
	DescriptionEn    *string `json:"description_en"`
	Status           *string `json:"status"`
	Body             *string `json:"body"`
	BodyKz           *string `json:"body_kz"`
	BodyEn           *string `json:"body_en"`
	CountOfQuestions *int    `json:"count_of_questions"`

	Image *Image `json:"image"`
}

type CreateContestPrize struct {
	ContestID uint   `json:"contest_id"`
	Number    int    `json:"number"`
	PrizeName string `json:"prize_name"`

	Image *Image `json:"image"`
}

type UpdateContestPrize struct {
	ContestPrizeID uint    `json:"contest_prize_id"`
	Number         *int    `json:"number"`
	PrizeName      *string `json:"prize_name"`

	Image *Image `json:"image"`
}
