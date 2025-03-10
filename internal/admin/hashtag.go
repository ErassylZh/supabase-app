package admin

type UpdateHashtag struct {
	HashtagID uint   `json:"hashtag_id"`
	Name      string `json:"name"`
	NameRu    string `json:"name_ru"`
	NameKz    string `json:"name_kz"`
	IsVisible bool   `json:"is_visible"`
	Image     *Image `json:"image"`
}

type CreateHashtag struct {
	Name      string `json:"name"`
	NameRu    string `json:"name_ru"`
	NameKz    string `json:"name_kz"`
	IsVisible bool   `json:"is_visible"`
	Image     *Image `json:"image"`
}

type AddHashtag struct {
	PostID    uint `json:"post_id"`
	HashtagID uint `json:"hashtag_id"`
}

type DeleteHashtagPost struct {
	PostID    uint `json:"post_id"`
	HashtagID uint `json:"hashtag_id"`
}
