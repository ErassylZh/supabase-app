package airtable

import "time"

type BaseResponse[T any] struct {
	Records []BaseObject[T] `json:"records"`
	Offset  *string         `json:"offset"`
}

type BaseObject[T any] struct {
	Id          string    `json:"id"`
	CreatedTime time.Time `json:"createdTime"`
	Fields      *T        `json:"fields"`
}

type ProductListResponse struct {
	Count         int      `json:"count"`
	Point         int      `json:"point"`
	Sapphire      int      `json:"sapphire"`
	Description   string   `json:"description"`
	DescriptionKz string   `json:"description_kz"`
	DescriptionEn string   `json:"description_en"`
	Title         string   `json:"title"`
	TitleEn       string   `json:"title_en"`
	TitleKz       string   `json:"title_kz"`
	Image         []Image  `json:"image"`
	Sku           string   `json:"sku"`
	ProductType   string   `json:"product_type"`
	SellType      string   `json:"sell_type"`
	Status        string   `json:"status"`
	Offer         string   `json:"offer"`
	OfferKz       string   `json:"offer_kz"`
	OfferEn       string   `json:"offer_en"`
	TagName       []string `json:"tag_name"`
	Discount      string   `json:"discount"`
	DiscountKz    string   `json:"discount_kz"`
	DiscountEn    string   `json:"discount_en"`
	Contacts      string   `json:"contacts"`
	ContactsKz    string   `json:"contacts_kz"`
	ContactsEn    string   `json:"contacts_en"`
}

type Image struct {
	FileName string `json:"filename"`
	File     string `json:"file"`
}

type Post struct {
	Title            string   `json:"title"`
	Company          *string  `json:"company"`
	Code             string   `json:"code"`
	Body             *string  `json:"body"`
	Status           *string  `json:"status"`
	RatingStatus     *string  `json:"rating_status"`
	Description      *string  `json:"description"`
	Image            []Image  `json:"image"`
	QuizTime         *int     `json:"quiz_time"`
	ReadTime         *int     `json:"read_time"`
	Point            *int     `json:"point"`
	Logo             []Image  `json:"logo"`
	Language         *string  `json:"language"`
	HashtagName      []string `json:"hashtag_name"`
	Uuid             string   `json:"uuid"`
	Hashtags         []string `json:"hashtags"`
	CollectionName   []string `json:"collection_name"`
	Sapphire         int      `json:"sapphire"`
	ShortDescription *string  `json:"short_description"`
	PostOrder        *int     `json:"order"`
}

type Stories struct {
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
	Order     int       `json:"order"`
	Image     []Image   `json:"image"`
	Uuid      string    `json:"uuid"`
	Icon      *[]Image  `json:"icon"`
}

type Hashtag struct {
	Name      string   `json:"name"`
	NameRu    string   `json:"name_ru"`
	NameKz    string   `json:"name_kz"`
	IsVisible bool     `json:"is_visible"`
	Image     *[]Image `json:"image"`
}

type ProductTag struct {
	Name   string   `json:"name"`
	NameRu string   `json:"name_ru"`
	NameKz string   `json:"name_kz"`
	Image  *[]Image `json:"image"`
}

type Collection struct {
	Name             string   `json:"name"`
	NameRu           string   `json:"name_ru"`
	NameKz           string   `json:"name_kz"`
	Image            *[]Image `json:"image"`
	ImageRu          *[]Image `json:"image_ru"`
	ImageKz          *[]Image `json:"image_kz"`
	IsRecommendation bool     `json:"is_recommendation"`
}

type Contest struct {
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	IsActive                 bool
	Code                     string `json:"uuid"`
	ConsolationPrizeSapphire int
}

type ContestBook struct {
	Body             string
	Title            string
	Description      string   `gorm:"column:description" json:"description"`
	DayNumber        int      `gorm:"column:day_number" json:"day_number"`
	ContestCoins     int      `gorm:"column:contest_coins" json:"contest_coins"`
	Point            int      `gorm:"column:point" json:"point"`
	CountOfQuestions int      `gorm:"column:count_of_questions"  json:"count_of_questions"`
	Image            *[]Image `json:"image"`
	Code             string   `json:"code"`
}

type ContestPrize struct {
	ContestPrizeID uint     `gorm:"primaryKey;column:contest_prize_id" json:"contest_prize_id"`
	ContestID      uint     `gorm:"column:contest_id" json:"contest_id"`
	Number         int      `gorm:"column:number" json:"number"`
	PrizeName      string   `gorm:"column:prize_name" json:"prize_name"`
	PhotoPath      string   `gorm:"column:photo_path" json:"photo_path"`
	Image          *[]Image `json:"image"`
	Code           string   `json:"code"`
}
