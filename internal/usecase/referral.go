package usecase

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
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
	return u.referralCode.GetByUserId(ctx, userID)
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
