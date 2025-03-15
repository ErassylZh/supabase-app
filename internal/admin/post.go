package admin

type CreatePost struct {
	Title            string  `json:"title"`
	Company          *string `json:"company"`
	Language         string  `json:"language"`
	Description      *string `json:"description"`
	Body             string  `json:"body"`
	ReadTime         *int    `json:"read_time"`
	RatingStatus     *string `json:"rating_status"`
	Status           string  `gorm:"column:status" json:"status"`
	Point            *int    `json:"point"`
	QuizTime         *int    `json:"quiz_time"`
	Sapphire         int     `json:"sapphire"`
	ShortDescription *string `json:"short_description"`
	Logo             *Image  `json:"logo"`
	Image            *Image  `json:"image"`
}

type UpdatePost struct {
	PostID           uint    `json:"post_id"`
	Title            *string `json:"title"`
	Company          *string `json:"company"`
	Language         *string `json:"language"`
	Description      *string `json:"description"`
	Status           *string `json:"status"`
	Body             *string `json:"body"`
	ReadTime         *int    `json:"read_time"`
	RatingStatus     *string `json:"rating_status"`
	Point            *int    `json:"point"`
	QuizTime         *int    `json:"quiz_time"`
	Sapphire         *int    `json:"sapphire"`
	ShortDescription *string `json:"short_description"`
	Logo             *Image  `json:"logo"`
	Image            *Image  `json:"image"`
}
