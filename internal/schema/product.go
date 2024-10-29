package schema

type ProductBuyRequest struct {
	ProductId uint   `json:"product_id"`
	UserId    string `json:"-"`
}
