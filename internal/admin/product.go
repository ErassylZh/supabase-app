package admin

type CreateProduct struct {
	Title         string `json:"title"`
	TitleEn       string `json:"title_en"`
	TitleKz       string `json:"title_kz"`
	Description   string `json:"description"`
	DescriptionKz string `json:"description_kz"`
	DescriptionEn string `json:"description_en"`
	Count         int    `json:"count"`
	Point         int    `gorm:"column:point" json:"point"`
	Sapphire      int    `gorm:"column:sapphire" json:"sapphire"`
	Sku           string `gorm:"column:sku" json:"sku"`
	ProductType   string `gorm:"column:product_type" json:"product_type"`
	Status        string `gorm:"column:status" json:"status"`
	SellType      string `gorm:"column:sell_type" json:"sell_type"`
	Offer         string `gorm:"column:offer" json:"offer"`
	OfferKz       string `gorm:"column:offer_kz" json:"offer_kz"`
	OfferEn       string `gorm:"column:offer_en" json:"offer_en"`
	Discount      string `gorm:"column:discount" json:"discount"`
	DiscountKz    string `gorm:"column:discount_kz" json:"discount_kz"`
	DiscountEn    string `gorm:"column:discount_en" json:"discount_en"`
	Contacts      string `gorm:"column:contacts" json:"contacts"`
	ContactsEn    string `gorm:"column:contacts_en" json:"contacts_en"`
	ContactsKz    string `gorm:"column:contacts_kz" json:"contacts_kz"`
	Logo          *Image `json:"logo"`
}

type UpdateProduct struct {
	ProductID     uint   `json:"product_id"`
	Title         string `json:"title"`
	TitleEn       string `json:"title_en"`
	TitleKz       string `json:"title_kz"`
	Description   string `json:"description"`
	DescriptionKz string `json:"description_kz"`
	DescriptionEn string `json:"description_en"`
	Count         int    `json:"count"`
	Point         int    `gorm:"column:point" json:"point"`
	Sapphire      int    `gorm:"column:sapphire" json:"sapphire"`
	Sku           string `gorm:"column:sku" json:"sku"`
	ProductType   string `gorm:"column:product_type" json:"product_type"`
	SellType      string `gorm:"column:sell_type" json:"sell_type"`
	Offer         string `gorm:"column:offer" json:"offer"`
	OfferKz       string `gorm:"column:offer_kz" json:"offer_kz"`
	OfferEn       string `gorm:"column:offer_en" json:"offer_en"`
	Discount      string `gorm:"column:discount" json:"discount"`
	DiscountKz    string `gorm:"column:discount_kz" json:"discount_kz"`
	DiscountEn    string `gorm:"column:discount_en" json:"discount_en"`
	Contacts      string `gorm:"column:contacts" json:"contacts"`
	ContactsEn    string `gorm:"column:contacts_en" json:"contacts_en"`
	ContactsKz    string `gorm:"column:contacts_kz" json:"contacts_kz"`
	Logo          *Image `json:"logo"`
}

type CreateProductTag struct {
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	NameKz string `json:"name_kz"`
}

type UpdateProductTag struct {
	ProductTagID uint    `json:"product_tag_id"`
	Name         *string `json:"name"`
	NameEn       *string `json:"name_en"`
	NameKz       *string `json:"name_kz"`
}

type AddProductProductTag struct {
	ProductID    uint `json:"product_id"`
	ProductTagID uint `json:"product_tag_id"`
}

type DeleteProductProductTag struct {
	ProductID    uint `json:"product_id"`
	ProductTagID uint `json:"product_tag_id"`
}
