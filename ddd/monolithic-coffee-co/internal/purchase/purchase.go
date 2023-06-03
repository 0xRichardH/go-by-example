package purchase

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/store"
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"

	"time"
)

type Purchase struct {
	id                 uuid.UUID
	total              money.Money
	timeOfPurchase     time.Time
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	PaymentMethod      payment.Means
	CardToken          *string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) <= 0 {
		return ErrInvalidPurchase
	}

	p.total = *money.New(0, "USD")
	for _, product := range p.ProductsToPurchase {
		newTotal, err := p.total.Add(&product.BasePrice)
		if err != nil {
			return fmt.Errorf("failed to calculate total: %w", err)
		}
		p.total = *newTotal
	}
	if p.total.IsZero() {
		return ErrInvalidPurchaseAmount
	}

	p.id = uuid.New()
	p.timeOfPurchase = time.Now()

	return nil
}
