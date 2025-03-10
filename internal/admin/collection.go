package admin

type UpdateCollection struct {
	CollectionID     uint   `json:"collection_id"`
	Name             string `json:"name"`
	NameRu           string `json:"name_ru"`
	NameKz           string `json:"name_kz"`
	IsRecommendation bool   `json:"is_recommendation"`
	Image            *Image `json:"image"`
	ImageRu          *Image `json:"image_ru"`
	ImageKz          *Image `json:"image_kz"`
}

type CreateCollection struct {
	Name             string `json:"name"`
	NameRu           string `json:"name_ru"`
	NameKz           string `json:"name_kz"`
	IsRecommendation bool   `json:"is_recommendation"`
	Image            *Image `json:"image"`
	ImageRu          *Image `json:"image_ru"`
	ImageKz          *Image `json:"image_kz"`
}

type AddCollection struct {
	PostID       uint `json:"post_id"`
	CollectionID uint `json:"collection_id"`
}

type DeleteCollectionPost struct {
	PostID       uint `json:"post_id"`
	CollectionID uint `json:"collection_id"`
}
