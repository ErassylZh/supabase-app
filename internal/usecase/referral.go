package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math/big"
	"time"
	"work-project/internal/model"
	"work-project/internal/repository"
	"work-project/internal/service"
)

type Referral interface {
	GetReferralCodeByUser(ctx context.Context, userID string) (model.ReferralCode, error)
	AcceptReferralCode(ctx context.Context, userID string, referralCode string) (model.ReferralCode, error)
}

type ReferralUsecase struct {
	referralCodeRepository  repository.ReferralCode
	referralRepository      repository.Referral
	pushNotificationService service.PushNotification
	balanceService          service.Balance
}

func NewReferralUsecase(referralCode repository.ReferralCode, referral repository.Referral, pushNotification service.PushNotification, balance service.Balance) *ReferralUsecase {
	return &ReferralUsecase{referralCodeRepository: referralCode, referralRepository: referral, pushNotificationService: pushNotification, balanceService: balance}
}

func (u *ReferralUsecase) GetReferralCodeByUser(ctx context.Context, userID string) (model.ReferralCode, error) {
	referralCode, err := u.referralCodeRepository.GetByUserId(ctx, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return referralCode, err
	}
	if err == nil {
		return referralCode, nil
	}
	code, err := u.generateReferralCode(referralCodeLength)
	if err != nil {
		return model.ReferralCode{}, err
	}
	referralCode = model.ReferralCode{
		UserID:       userID,
		ReferralCode: code,
		CreatedAt:    time.Now(),
	}
	err = u.referralCodeRepository.Create(ctx, referralCode)
	if err != nil {
		return model.ReferralCode{}, err
	}
	return referralCode, nil
}

func (u *ReferralUsecase) AcceptReferralCode(ctx context.Context, userID string, code string) (model.ReferralCode, error) {
	referralCode, err := u.referralCodeRepository.GetByReferralCode(ctx, code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ReferralCode{}, fmt.Errorf("referral code not found")
		}
		return model.ReferralCode{}, err
	}
	_, err = u.referralRepository.GetByInvitedUserId(ctx, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.ReferralCode{}, err
	}
	if err == nil {
		return model.ReferralCode{}, fmt.Errorf("user already accept referral code ")
	}

	referral := model.Referral{
		UserID:        referralCode.UserID,
		InvitedUserID: userID,
		ReferralCode:  code,
		CreatedAt:     time.Now(),
	}
	err = u.referralRepository.Create(ctx, referral)
	if err != nil {
		return model.ReferralCode{}, err
	}
	//todo добавить отправвку пуша и добавление бонусов двум юзерам
	//err = u.pushNotification.Send(ctx, "some text", "some header", "some token", nil, nil)

	//todo потом возможно поменяем койны и сапфиры которые даем за реферал
	_, err = u.balanceService.CreateTransaction(ctx, referralCode.UserID, model.Transaction{
		Coins:           10,
		Sapphires:       0,
		UserId:          referralCode.UserID,
		CreatedAt:       time.Now(),
		TransactionType: string(model.TRANSACTION_TYPE_INCOME),
		Reason:          string(model.TRANSACTION_REASON_REFERRAL),
	})
	if err != nil {
		return model.ReferralCode{}, err
	}

	_, err = u.balanceService.CreateTransaction(ctx, userID, model.Transaction{
		Coins:           10,
		Sapphires:       0,
		UserId:          userID,
		CreatedAt:       time.Now(),
		TransactionType: string(model.TRANSACTION_TYPE_INCOME),
		Reason:          string(model.TRANSACTION_REASON_REFERRAL),
	})
	if err != nil {
		return model.ReferralCode{}, err
	}
	return model.ReferralCode{}, nil
}

const referralCodeLength = 10
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (u *ReferralUsecase) generateReferralCode(length int) (string, error) {
	code := make([]byte, length)
	for i := range code {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[randomInt.Int64()]
	}
	return string(code), nil
}
