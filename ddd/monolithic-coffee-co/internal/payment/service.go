package payment

import (
	"coffeeco/internal/payment/service"
	"context"

	"github.com/Rhymond/go-money"
)

type CardChargeService interface {
	ChargeCard(ctx context.Context, amount money.Money, cardToken string) error
}

var _ CardChargeService = (*service.StripeService)(nil)

var NewStripeService = service.NewStripeService
