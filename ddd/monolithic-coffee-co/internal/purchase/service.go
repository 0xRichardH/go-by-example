package purchase

import (
	"coffeeco/internal/loyalty"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
	"context"
	"fmt"
)

type Service struct {
	cardChargeService payment.CardChargeService
	storeService      store.Service
	purchaseRepo      Repository
}

func (s Service) CompletePurchase(ctx context.Context, purchase *Purchase, coffeeBuxCard *loyalty.CoffeeBux) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return fmt.Errorf("failed to validate and enrich purchase: %w", err)
	}

	if err := s.calculateStoreSpecificDiscount(ctx, purchase); err != nil {
		return fmt.Errorf("failed to calculate store specific discount: %w", err)
	}

	switch purchase.PaymentMethod {
	case payment.MEANS_CARD:
		if err := s.cardChargeService.ChargeCard(ctx, purchase.total, *purchase.CardToken); err != nil {
			return fmt.Errorf("failed to charge card: %w", err)
		}
	case payment.MEANS_CASH:
		// TODO: chage cash
	case payment.MEANS_COFFEEBUX:
		if err := coffeeBuxCard.Pay(ctx, purchase.ProductsToPurchase); err != nil {
			return fmt.Errorf("failed to pay coffeeBUX: %w", err)
		}
	default:
		return ErrUnknownPaymentMethod
	}

	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return fmt.Errorf("failed to store purchase: %w", err)
	}

	if coffeeBuxCard != nil {
		coffeeBuxCard.AddStamp()
	}

	return nil
}

func (s *Service) calculateStoreSpecificDiscount(ctx context.Context, purchase *Purchase) error {
	discount, err := s.storeService.GetStoreSpecificDiscount(ctx, purchase.Store.ID)
	if err != nil && err != store.ErrNoDiscount {
		return fmt.Errorf("failed to get store discount: %w", err)
	}

	purchasePrice := purchase.total
	if discount > 0 {
		purchase.total = *purchasePrice.Multiply(int64(100 - discount))
	}

	return nil
}
