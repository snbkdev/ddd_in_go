package store

import (
	"github.com/google/uuid"
	"context"
	"DDD/coffeeco/internal"
)

type Store struct {
	ID uuid.UUID
	Location string
	ProductsForSale []internal.Product
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	dis, err := s.repo.GetStoreDiscount(ctx, storeID)
	if err != nil {
		return 0, err
	}

	return float32(dis), nil
}

type Service struct {
	repo Repository
}