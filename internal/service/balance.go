package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Balance interface {
	GetByUserId(ctx context.Context, userId string) (model.Balance, error)
	GetTransactionHistory(ctx context.Context, userId string) ([]model.Transaction, error)
	CreateTransaction(ctx context.Context, userId string, transaction model.Transaction) (model.Transaction, error)
}

type BalanceService struct {
	balanceRepo     repository.Balance
	transactionRepo repository.Transaction
}

func NewBalanceService(balanceRepo repository.Balance, transactionRepo repository.Transaction) *BalanceService {
	return &BalanceService{balanceRepo: balanceRepo, transactionRepo: transactionRepo}
}

func (s *BalanceService) GetByUserId(ctx context.Context, userId string) (model.Balance, error) {
	return s.balanceRepo.GetByUserID(ctx, userId)
}

func (s *BalanceService) GetTransactionHistory(ctx context.Context, userId string) ([]model.Transaction, error) {
	return s.transactionRepo.GetAllByUserID(ctx, userId)
}

func (s *BalanceService) CreateTransaction(ctx context.Context, userId string, transaction model.Transaction) (model.Transaction, error) {
	balance, err := s.balanceRepo.GetByUserID(ctx, userId)
	if err != nil {
		return model.Transaction{}, err
	}
	balance.Coins += transaction.Coins
	balance.Sapphires += transaction.Sapphires
	err = s.balanceRepo.Update(ctx, balance)
	if err != nil {
		return model.Transaction{}, err
	}

	err = s.transactionRepo.Create(ctx, transaction)
	if err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}
