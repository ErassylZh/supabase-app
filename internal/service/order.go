package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Order interface {
	CreateOrder(ctx context.Context, userId string, products []model.Product) (model.Order, error)
	GetOrders(ctx context.Context, id string) ([]model.Order, error)
}

type OrderService struct {
	orderRepo         repository.Order
	orderProductsRepo repository.OrderProduct
}

func NewOrderService(orderRepo repository.Order, orderProductsRepo repository.OrderProduct) *OrderService {
	return &OrderService{orderRepo: orderRepo, orderProductsRepo: orderProductsRepo}
}

func (s *OrderService) CreateOrder(ctx context.Context, userId string, products []model.Product) (model.Order, error) {
	order, err := s.orderRepo.Create(ctx, model.Order{
		BuyerId: &userId,
	})
	if err != nil {
		return model.Order{}, err
	}

	orderProducts := make([]model.OrderProduct, len(products))
	for i := range products {
		orderProducts[i] = model.OrderProduct{
			OrderID:     order.OrderId,
			ProductID:   products[i].ProductID,
			Quantity:    1,
			Coins:       products[i].Point,
			Sku:         products[i].Sku,
			ProductType: products[i].ProductType,
		}
	}
	orderProducts, err = s.orderProductsRepo.Create(ctx, orderProducts)
	if err != nil {
		return model.Order{}, err
	}
	order.OrderProducts = orderProducts
	return order, nil
}

func (s *OrderService) GetOrders(ctx context.Context, userId string) ([]model.Order, error) {
	return s.orderRepo.GetByBuyerId(ctx, userId)

}
