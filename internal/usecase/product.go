package usecase

import (
	"context"
	"errors"
	"strconv"
	"work-project/internal/model"
	"work-project/internal/schema"
	"work-project/internal/service"
)

type Product interface {
	Buy(ctx context.Context, request schema.ProductBuyRequest) error
}

type ProductUsecase struct {
	productService service.Product
	balanceService service.Balance
	orderService   service.Order
	emailSender    service.EmailSender
	userService    service.User
}

func NewProductUsecase(
	productService service.Product,
	balanceService service.Balance,
	orderService service.Order,
	emailSender service.EmailSender,
	userService service.User,
) *ProductUsecase {
	return &ProductUsecase{productService: productService, balanceService: balanceService, orderService: orderService, emailSender: emailSender, userService: userService}
}

func (u *ProductUsecase) Buy(ctx context.Context, request schema.ProductBuyRequest) error {
	user, err := u.userService.GetById(ctx, request.UserId)
	if err != nil {
		return err
	}

	product, err := u.productService.GetProduct(ctx, request.ProductId)
	if err != nil {
		return err
	}

	balance, err := u.balanceService.GetByUserId(ctx, request.UserId)
	if err != nil {
		return err
	}
	if balance.Coins < product.Point {
		return errors.New("not enough money :(")
	}

	_, err = u.balanceService.CreateTransaction(ctx, request.UserId, model.Transaction{
		UserId:            request.UserId,
		TransactionType:   string(model.TRANSACTION_TYPE_SPEND),
		Coins:             product.Point,
		TransactionReason: string(model.TRANSACTION_REASON_STORE),
	})

	order, err := u.orderService.CreateOrder(ctx, request.UserId, []model.Product{product})
	if err != nil {
		return err
	}

	if product.ProductType == string(model.PRODUCT_TYPE_VIRTUAL) {
		err = u.emailSender.Send(ctx, schema.Message{
			Subject:     "Order #" + strconv.Itoa(int(order.OrderId)),
			Body:        "SPASSSSSSSSSSIBA ZA POKUPKU BRATES'",
			To:          []string{user.Email},
			Attachments: nil,
		})
		if err != nil {
			return err
		}

	}
	return nil
}
