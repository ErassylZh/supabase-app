package service

import (
	"context"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type PrivacyTerms interface {
	GetAll(ctx context.Context) ([]model.PrivacyTerms, error)
}

type PrivacyTermsService struct {
	privacyTermsRepo repository.PrivacyTerms
}

func NewPrivacyTermsService(privacyTermsRepo repository.PrivacyTerms) *PrivacyTermsService {
	return &PrivacyTermsService{privacyTermsRepo: privacyTermsRepo}
}

func (s *PrivacyTermsService) GetAll(ctx context.Context) ([]model.PrivacyTerms, error) {
	return s.privacyTermsRepo.GetAll(ctx)
}
