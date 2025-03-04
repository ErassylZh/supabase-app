package admin

type UpdateCollection struct {
	CollectionID     uint    `json:"collection_id"`
	Name             string  `json:"name"`
	NameRu           string  `json:"name_ru"`
	NameKz           string  `json:"name_kz"`
	IsRecommendation bool    `json:"is_recommendation"`
	ImageBase64      *string `json:"image_base64"`
	ImageKzBase64    *string `json:"image_kz_base64"`
	ImageRuBase64    *string `json:"image_ru_base64"`
}

type CreateCollection struct {
	Name             string  `json:"name"`
	NameRu           string  `json:"name_ru"`
	NameKz           string  `json:"name_kz"`
	IsRecommendation bool    `json:"is_recommendation"`
	ImageBase64      *string `json:"image_base64"`
	ImageKzBase64    *string `json:"image_kz_base64"`
	ImageRuBase64    *string `json:"image_ru_base64"`
}
