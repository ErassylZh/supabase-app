package airtable

import "time"

type BaseResponse[T any] struct {
	Records []BaseObject[T] `json:"records"`
}

type BaseObject[T any] struct {
	Id          string    `json:"id"`
	CreatedTime time.Time `json:"createdTime"`
	Fields      []T       `json:"fields"`
}

type ProductListResponse struct {
	Count       int     `json:"count"`
	Point       int     `json:"point"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Image       []Image `json:"image"`
}

type Image struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	FileName string `json:"filename"`
	Type     string `json:"type"`
}

type Post struct {
	Title        string   `json:"title"`
	Company      string   `json:"company"`
	Body         string   `json:"body"`
	Status       string   `json:"status"`
	RatingStatus string   `json:"rating_status"`
	Description  string   `json:"description"`
	Image        []Image  `json:"image"`
	QuizTime     int      `json:"quiz_time"`
	ReadTime     int      `json:"read_time"`
	Point        int      `json:"point"`
	Logo         []Image  `json:"logo"`
	Language     string   `json:"language"`
	HashtagName  []string `json:"hashtag_name"`
}
