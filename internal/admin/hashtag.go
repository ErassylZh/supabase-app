package admin

type UpdateHashtag struct {
	HashtagID   uint    `json:"hashtag_id"`
	Name        string  `json:"name"`
	NameRu      string  `json:"name_ru"`
	NameKz      string  `json:"name_kz"`
	IsVisible   bool    `json:"is_visible"`
	ImageBase64 *string `json:"image_base64"`
}

type CreateHashtag struct {
	Name        string  `json:"name"`
	NameRu      string  `json:"name_ru"`
	NameKz      string  `json:"name_kz"`
	IsVisible   bool    `json:"is_visible"`
	ImageBase64 *string `json:"image_base64"`
}
