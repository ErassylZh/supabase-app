package admin

import "time"

type CreateStories struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Title     string    `json:"title"`
	Icon      *Image    `json:"icon"`
}

type UpdateStories struct {
	StoriesID uint       `json:"stories_id"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Title     *string    `json:"title"`
	Icon      *Image     `json:"icon"`
}

type CreateStoryPage struct {
	StoriesId uint   `json:"stories_id"`
	Text      string `json:"text"`
	PageOrder int    `json:"page_order"`
	Icon      *Image `json:"icon"`
}

type UpdateStoryPage struct {
	StoryPageID uint    `json:"story_page_id"`
	Text        *string `json:"text"`
	PageOrder   *int    `json:"page_order"`
	Icon        *Image  `json:"icon"`
}
