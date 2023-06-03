package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	discount, err := s.repo.GetStoreDiscount(ctx, storeID)
	if err != nil {
		return 0, fmt.Errorf("failed to get store discount: %w", err)
	}
	return float32(discount), nil
}
