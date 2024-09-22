package schema

type CollectionListResponse struct {
	CollectionID     uint    `json:"collection_id"`
	Name             string  `json:"name"`
	NameRu           string  `json:"name_ru"`
	NameKz           string  `json:"name_kz"`
	ImagePath        *string `json:"image_path"`
	ImagePathKz      *string `json:"image_path_kz"`
	ImagePathRu      *string `json:"image_path_ru"`
	IsRecommendation bool    `json:"is_recommendation"`

	Posts []PostResponse `json:"posts"`
}
