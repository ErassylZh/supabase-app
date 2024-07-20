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
)

type Referral interface {
	GetReferralCodeByUser(ctx context.Context, userID string) (model.ReferralCode, error)
	AcceptReferralCode(ctx context.Context, userID string, referralCode string) (model.ReferralCode, error)
}

type ReferralUsecase struct {
	referralCode repository.ReferralCode
	referral     repository.Referral
}

func NewReferralUsecase(referralCode repository.ReferralCode, referral repository.Referral) *ReferralUsecase {
	return &ReferralUsecase{
		referralCode: referralCode,
		referral:     referral,
	}
}

func (u *ReferralUsecase) GetReferralCodeByUser(ctx context.Context, userID string) (model.ReferralCode, error) {
	referralCode, err := u.referralCode.GetByUserId(ctx, userID)
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
	err = u.referralCode.Create(ctx, referralCode)
	if err != nil {
		return model.ReferralCode{}, err
	}
	return referralCode, nil
}

func (u *ReferralUsecase) AcceptReferralCode(ctx context.Context, userID string, code string) (model.ReferralCode, error) {
	referralCode, err := u.referralCode.GetByReferralCode(ctx, code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ReferralCode{}, fmt.Errorf("referral code not found")
		}
		return model.ReferralCode{}, err
	}

	referral := model.Referral{
		UserID:        referralCode.UserID,
		InvitedUserID: userID,
		ReferralCode:  code,
		CreatedAt:     time.Now(),
	}
	err = u.referral.Create(ctx, referral)
	if err != nil {
		return model.ReferralCode{}, err
	}
	//todo добавить отправвку пуша и добавление бонусов двум юзерам
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
